init:
	pip3 install -r xgb-python-unixsocket-rawbytes/requirements.txt

bench: clean
	go test -bench=. -benchtime=10s -benchmem -cpu=1 ./...

clean:
	-rm *sc
