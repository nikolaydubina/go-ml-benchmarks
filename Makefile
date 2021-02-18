PWD := $(shell echo $$PWD)
MY_INSTALL_DIR := $(shell echo $$HOME/.local)
XGBOOST_ROOT := $(PWD)/third_party/xgboost

init:
	sudo apt-get install python3-pip
	pip3 install jupyter
	cd go-client; go mod download && go generate ./...
	cd third_party/grpc; git submodule update --init --recursive
	cd third_party/xgboost; git submodule update --init --recursive
	cd third_party/FlameGraph; git submodule update --init --recursive
	sudo sysctl -w kernel.perf_event_paranoid=0
	sudo apt-get install linux-tools-common linux-tools-5.4.0-1037-aws

leaves:
	mkdir -p docs/profiles/leaves
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		MODEL_PATH=$(PWD)/data/models/titanic_v090.xgb \
		go test -bench=BenchmarkXGB_Go_GoFeatureProcessing_GoLeaves* -benchtime=10s -cpu=1 -cpuprofile $(PWD)/docs/profiles/leaves/cpu.profile ./main | tee -a $(PWD)/docs/bench.out

cgo:
	mkdir -p docs/profiles/cgo; cd docs/profiles/cgo; perf record -F 99 -a -g -- sleep 30 &
	cd cgo-version; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		CGO_CFLAGS="-I$(XGBOOST_ROOT)/include -I$(XGBOOST_ROOT)/dmlc-core/include -I$(XGBOOST_ROOT)/rabit/include" \
		CGO_LDFLAGS="-L$(XGBOOST_ROOT)/lib -L$(MY_INSTALL_DIR)/lib -lxgboost -ldmlc -lstdc++ -lm -fopenmp" \
		go test -bench=BenchmarkXGB_CGo* -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	sleep 20

uds:
	pip3 install -r bench-gofeatureprocessing-uds-raw-python-xgb/requirements.txt
	mkdir -p docs/profiles/uds; cd docs/profiles/uds; perf record -F 99 -a -g -- sleep 20 &
	cd bench-gofeatureprocessing-uds-raw-python-xgb; \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		SOCKET_PATH=$(PWD)/sc \
		python3 main.py & echo "$$!" > pids
	sleep 5
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		SOCKET_PATH=$(PWD)/sc \
		go test -bench=BenchmarkXGB_Go_GoFeatureProcessing_UDS_RawBytes_Python_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	cd bench-gofeatureprocessing-uds-raw-python-xgb; kill -9 $$(cat pids); rm pids
	sleep 10

rest:
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; pip3 install -r requirements.txt
	mkdir -p docs/profiles/rest; cd docs/profiles/rest; perf record -F 99 -a -g -- sleep 20 &
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		gunicorn --workers=3 --threads=2 --bind=0.0.0.0:1024 wsgi:app & 
	sleep 5
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_Go_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	-pkill -f gunicorn
	sleep 10

init-grpc-go:
	sudo apt-get install -y protobuf-compiler
	export GO111MODULE=on
	export PATH="$$PATH:$(go env GOPATH)/bin"
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
	mkdir -p docs/profiles/grpc-python-sklearn; cd docs/profiles/grpc-python-sklearn; perf record -F 99 -a -g -- sleep 20 &
	cd bench-uds-grpc-python-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		python3 main.py &
	sleep 5
	cd go-client; \
		GO111MODULE=on \
		go test -bench=BenchmarkXGB_Go_UDS_gRPC_Python_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	-sudo pkill -f Python
	-sudo pkill -f python
	sleep 10

grpc-python-processed: init-grpc-go init-grpc-python
	mkdir -p docs/profiles/grpc-python-processed; cd docs/profiles/grpc-python-processed; perf record -F 99 -a -g -- sleep 20 &
	cd bench-uds-grpc-python-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		python3 main.py &
	sleep 5
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_Go_GoFeatureProcessing_UDS_gRPC_Python_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	-sudo pkill -f Python
	-sudo pkill -f python
	sleep 10

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
	mkdir -p docs/profiles/grpc-cpp; cd docs/profiles/grpc-cpp; perf record -F 99 -a -g -- sleep 20 &
	cd bench-uds-grpc-cpp-xgb; \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		./cmake/build/predictor &
	sleep 5
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_Go_GoFeatureProcessing_UDS_gRPC_CPP_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	pkill -f predictor
	sleep 10

bench: clean leaves cgo uds rest grpc-python-sklearn grpc-python-processed grpc-cpp

.PHONY: docs/profiles/*/perf.data
docs/profiles/*/perf.data:
	cd $(dir $@); \
		perf script > out.perf; \
		$(PWD)/third_party/FlameGraph/stackcollapse-perf.pl out.perf > out.folded; \
		$(PWD)/third_party/FlameGraph/flamegraph.pl out.folded > folded.svg;

docs: docs/profiles/*/perf.data
	cat docs/bench.out | grep Benchmark > docs/bench-clean.out

clean:
	jupyter nbconvert --clear-output --inplace notebooks/*.ipynb
	-rm -rf docs/profiles
	-sudo pkill -f Python
	-sudo pkill -f python
	-sudo pkill -f gunicorn
	-rm sc
	-rm docs/bench.out
	-cd bench-gofeatureprocessing-uds-raw-python-xgb; sudo kill -9 $$(cat pids); rm pids
	-sudo pkill -f predictor
	-rm -rf build
	-rm -rf bench-uds-grpc-cpp-xgb/cmake