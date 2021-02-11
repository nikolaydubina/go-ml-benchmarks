UNAME := $(shell uname)

init:
	cd go-client; go mod download && go generate ./...

leaves:
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_GoFeatureProcessing_GoLeaves -benchtime=10s -cpu=1 ./... | tee -a docs/bench.out

uds:
	pip3 install -r bench-uds-raw-python-xgb/requirements.txt
	cd bench-uds-raw-python-xgb; MODEL_PATH=../data/models/titanic.xgb SOCKET_PATH=../sc python3 main.py & echo "$$!" > pids
	sleep 3
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_GoFeatureProcessing_UDS_RawBytes_Python_XGB -benchtime=10s -cpu=1 ./... | tee -a docs/bench.out
	-kill -9 $$(cat pids)
	-rm pids sc

rest:
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; pip3 install -r requirements.txt
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; \
	    PREPROCESSOR_PATH=../data/models/titanic_preprocessor.sklearn \
	    MODEL_PATH=../data/models/titanic.xgb \
	    gunicorn --workers=3 --threads=2 --bind=0.0.0.0:80 wsgi:app & 
	sleep 7
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB -benchtime=10s -cpu=1 ./... | tee -a docs/bench.out
	-pkill -f gunicorn

grpc-python:
ifeq ($(UNAME), Darwin)
	brew install protobuf
endif
	pip3 install -r bench-grpc-python-sklearn-xgb/requirements.txt
	# apt install -y protobuf-compiler # for linux
	export GO111MODULE=on
	export PATH="$PATH:$(go env GOPATH)/bin"
	go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	cd bench-grpc-python-sklearn-xgb; python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. predictor.proto

bench: leaves uds rest grpc-python
	#cat docs/bench.out | grep Benchmark | column -t > docs/bench-clean.out
ifeq ($(UNAME), Darwin)
	brew install align
	cat docs/bench.out | grep Benchmark | column -t | align > docs/bench-clean.out
endif

clean:
	jupyter nbconvert --clear-output --inplace notebooks/*.ipynb
	-pkill -f gunicorn
	-rm docs/bench.out
