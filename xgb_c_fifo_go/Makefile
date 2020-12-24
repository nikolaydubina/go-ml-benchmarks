bench:
	$(shell mkdir -p profiles)
	go test -bench=. -benchtime=10s -benchmem -memprofile profiles/mem.profile -cpuprofile profiles/cpu.profile -cpu=1 ./...

clean:
	rm *.fifo
	rm -r profiles
