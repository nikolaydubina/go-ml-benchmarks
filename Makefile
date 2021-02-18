PWD := $(shell echo $$PWD)
MY_INSTALL_DIR := $(shell echo $$HOME/.local)

init:
	sudo apt-get install python3-pip
	pip3 install jupyter
	cd go-client; go mod download && go generate ./...
	cd third_party/grpc; git submodule update --init --recursive
	cd third_party/xgboost; git submodule update --init --recursive

leaves:
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		MODEL_PATH=$(PWD)/data/models/titanic_v090.xgb \
		go test -bench=BenchmarkXGB_GoFeatureProcessing_GoLeaves* -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out

uds:
	pip3 install -r bench-gofeatureprocessing-uds-raw-python-xgb/requirements.txt
	cd bench-gofeatureprocessing-uds-raw-python-xgb; \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		SOCKET_PATH=$(PWD)/sc \
		python3 main.py & echo "$$!" > pids
	sleep 3
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		SOCKET_PATH=$(PWD)/sc \
		go test -bench=BenchmarkXGB_GoFeatureProcessing_UDS_RawBytes_Python_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	cd bench-gofeatureprocessing-uds-raw-python-xgb; kill -9 $$(cat pids); rm pids

rest:
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; pip3 install -r requirements.txt
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		gunicorn --workers=3 --threads=2 --bind=0.0.0.0:1024 wsgi:app & 
	sleep 7
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	-pkill -f gunicorn

init-grpc-go:
	sudo apt-get install -y protobuf-compiler
	export GO111MODULE=on
	export PATH="$PATH:$(go env GOPATH)/bin"
	go get \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc \
		--go_out=go-client --go_opt=paths=source_relative \
 		--go-grpc_out=go-client --go-grpc_opt=paths=source_relative \
		proto/predictor.proto
	go get google.golang.org/grpc

init-grpc-python:
	pip3 install grpcio grpcio-tools
	python3 -m grpc_tools.protoc -I. --python_out=bench-uds-grpc-python-xgb --grpc_python_out=bench-uds-grpc-python-xgb proto/predictor.proto

grpc-python-sklearn: init-grpc-go init-grpc-python
	cd bench-uds-grpc-python-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		python3 main.py &
	sleep 3
	cd go-client; \
		GO111MODULE=on \
		go test -bench=BenchmarkXGB_UDS_gRPC_Python_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	-sudo pkill -f Python
	-sudo pkill -f python

grpc-python-processed: init-grpc-go init-grpc-python
	cd bench-uds-grpc-python-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		python3 main.py &
	sleep 3
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_GoFeatureProcessing_UDS_gRPC_Python_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	-sudo pkill -f Python
	-sudo pkill -f python

cpp-grpc-lib:
	mkdir -p $(MY_INSTALL_DIR)
	export PATH="$$PATH:$(MY_INSTALL_DIR)/bin"
	sudo apt-get install -y autoconf automake libtool pkg-config cmake
	cd third_party/grpc; \
		mkdir -p cmake/build; \
		cd cmake/build; \
		cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$(MY_INSTALL_DIR) ../..; \
		make -j 4; \
		make install

cpp-xgb-lib:
	cd third_party/xgboost; \
		mkdir -p build; \
		cd build; \
		cmake -DBUILD_STATIC_LIB=ON -DCMAKE_INSTALL_PREFIX=$(MY_INSTALL_DIR) ..; \
		make install -j 4

grpc-cpp-build: cpp-xgb-lib cpp-grpc-lib
	mkdir -p bench-uds-grpc-cpp-xgb/cmake/build
	cd bench-uds-grpc-cpp-xgb/cmake/build; \
		cmake -DCMAKE_PREFIX_PATH=$(MY_INSTALL_DIR) ../..; \
		make -j 4

grpc-cpp: grpc-cpp-build
	cd bench-uds-grpc-cpp-xgb; \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		./cmake/build/predictor &
	sleep 3
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_GoFeatureProcessing_UDS_gRPC_CPP_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	pkill -f predictor

bench: clean leaves uds rest grpc-python-sklearn grpc-python-processed grpc-cpp
	cat docs/bench.out | grep Benchmark > docs/bench-clean.out

clean:
	jupyter nbconvert --clear-output --inplace notebooks/*.ipynb
	-sudo pkill -f Python
	-sudo pkill -f python
	-sudo pkill -f gunicorn
	-rm sc
	-rm docs/bench.out
	-cd bench-gofeatureprocessing-uds-raw-python-xgb; sudo kill -9 $$(cat pids); rm pids
	-sudo pkill -f predictor
	-rm -rf build
	-rm -rf bench-uds-grpc-cpp-xgb/cmake