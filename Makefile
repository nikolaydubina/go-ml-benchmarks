init:
	cd go-client; go mod download
	cd go-client; go generate ./...	
	pip3 install -r python-raw-uds-xgb/requirements.txt
	pip3 install -r python-json-gunicorn-flask-sklearn-xgb/requirements.txt

bench: bench-xgb-leaves bench-python-raw-uds-xgb bench-python-json-gunicorn-flask-sklearn-xgb

bench-xgb-leaves:
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_GoFeatureProcessing_GoLeaves -benchtime=10s -cpu=1 ./...

bench-python-raw-uds-xgb:
	python3 python-raw-uds-xgb/main.py sc data/models/titanic.xgb & echo "$$!" > pids
	sleep 3
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Python_UDS_RawBytes_NewConnection -benchtime=10s -cpu=1 ./...
	-kill -9 $$(cat pids)
	-rm pids sc

bench-python-json-gunicorn-flask-sklearn-xgb:
	cd python-json-gunicorn-flask-sklearn-xgb; PREPROCESSOR_PATH=../data/models/titanic_preprocessor.sklearn MODEL_PATH=../data/models/titanic.xgb gunicorn --workers=3 --threads=2 --bind=0.0.0.0:80 wsgi:app &
	sleep 7
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Python_JSON_Gunicorn_Flask_sklearn_xgb -benchtime=10s -cpu=1 ./...
	-pkill -f gunicorn

clean:
	-pkill -f gunicorn