init:
	cd go-client; go mod download
	cd go-client; go generate ./...	
	pip3 install -r python-raw-uds-xgb/requirements.txt

bench: bench-python-raw-uds-xgb bench-xgb-leaves

bench-python-raw-uds-xgb: clean
	python3 python-raw-uds-xgb/main.py sc data/models/titanic.xgb & echo "$$!" > pids
	sleep 3
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Python_UDS_RawBytes_NewConnection -benchtime=10s -cpu=1 ./...

bench-xgb-leaves: clean
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Leaves.* -benchtime=10s -cpu=1 ./...

clean:
	-rm *sc
	-kill -9 $$(cat pids)
	-rm pids
