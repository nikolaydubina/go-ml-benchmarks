PWD := $(shell echo $$PWD)
UNAME := $(shell uname)

init:
	cd go-client; go mod download && go generate ./...

leaves:
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		MODEL_PATH=$(PWD)/data/models/titanic_v090.xgb \
		go test -bench=BenchmarkXGB_GoFeatureProcessing_GoLeaves -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out

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
		gunicorn --workers=3 --threads=2 --bind=0.0.0.0:80 wsgi:app & 
	sleep 7
	cd go-client; \
		GO111MODULE=on \
		PREPROCESSOR_PATH=$(PWD)/data/models/go-featureprocessor.json \
		go test -bench=BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out
	-pkill -f gunicorn

init-grpc-go:
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
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
	python3 -m grpc_tools.protoc -I. --python_out=bench-uds-grpc-python-sklearn-xgb --grpc_python_out=bench-uds-grpc-python-sklearn-xgb proto/predictor.proto

grpc-python: init-grpc-go init-grpc-python
	cd bench-uds-grpc-python-sklearn-xgb; \
		PREPROCESSOR_PATH=$(PWD)/data/models/titanic_preprocessor.sklearn \
		MODEL_PATH=$(PWD)/data/models/titanic.xgb \
		python3 main.py &
	sleep 3
	cd go-client; \
		GO111MODULE=on \
		go test -bench=BenchmarkXGB_UDS_gRPC_Python_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a $(PWD)/docs/bench.out	
	-pkill -f Python

bench: clean leaves uds rest grpc-python
	cat docs/bench.out | grep Benchmark > docs/bench-clean.out

clean:
	jupyter nbconvert --clear-output --inplace notebooks/*.ipynb
	-pkill -f Python
	-pkill -f gunicorn
	-rm sc
	-rm docs/bench.out
	-cd bench-gofeatureprocessing-uds-raw-python-xgb; kill -9 $$(cat pids); rm pids