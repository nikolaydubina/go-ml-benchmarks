init:
	cd go-client; go mod download
	cd go-client; go generate ./...	
	pip3 install -r python-xgb-uds-raw/requirements.txt

bench: bench-python-xgb-uds-raw

bench-python-xgb-uds-raw: clean
	python3 python-xgb-uds-raw/main.py sc data/models/titanic.xgb & echo "$$!" > pids
	sleep 3
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Python_UDS_RawBytes_NewConnection -benchtime=10s -benchmem -cpu=1 ./...

bench-xgb-leaves: clean
	PROJECT_PATH=$$PWD go test -bench=BenchmarkXGB_Leaves.* -benchtime=10s -benchmem -cpu=1 ./...

clean:
	-rm *sc
	-kill -9 $$(cat pids)
	-rm pids
