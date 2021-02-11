init:
	cd go-client; go mod download
	cd go-client; go generate ./...	
	pip3 install -r bench-uds-raw-python-xgb/requirements.txt
	pip3 install -r bench-http-json-python-gunicorn-flask-sklearn-xgb/requirements.txt

bench: bench-leaves
bench-leaves:
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_GoFeatureProcessing_GoLeaves -benchtime=10s -cpu=1 ./...

bench: uds 
uds:
	cd bench-uds-raw-python-xgb; MODEL_PATH=../data/models/titanic.xgb SOCKET_PATH=../sc python3 main.py & echo "$$!" > pids
	sleep 3
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_GoFeatureProcessing_UDS_RawBytes_Python_XGB -benchtime=10s -cpu=1 ./...
	-kill -9 $$(cat pids)
	-rm pids sc

bench: rest
rest:
	cd bench-http-json-python-gunicorn-flask-sklearn-xgb; PREPROCESSOR_PATH=../data/models/titanic_preprocessor.sklearn MODEL_PATH=../data/models/titanic.xgb gunicorn --workers=3 --threads=2 --bind=0.0.0.0:80 wsgi:app &
	sleep 7
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB -benchtime=10s -cpu=1 ./...
	-pkill -f gunicorn

clean:
	-pkill -f gunicorn
