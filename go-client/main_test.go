package main_test

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

// predictRawBytes uses raw floats encoding
func predictRawBytes(w io.Writer, r io.Reader, features []float32) (float32, error) {
	if err := binary.Write(w, binary.LittleEndian, features); err != nil {
		return 0, err
	}

	var prediction float32
	if err := binary.Read(r, binary.LittleEndian, &prediction); err != nil {
		return 0, err
	}

	return prediction, nil
}

func benchmarkUNIXSocketRawBytesNewConn(b *testing.B, socketpathin string) {
	addr, err := net.ResolveUnixAddr("unix", socketpathin)
	if err != nil {
		panic(err)
	}

	features := []float32{72.0, 72.0, 69.0, 60.0, 7.0, 0, 0, 0, 1, 1, 0, 1, 0}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c, err := net.DialUnix("unix", nil, addr)
		if err != nil {
			panic(err)
		}
		result, err := predictRawBytes(c, c, features)
		if result == 0 || err != nil {
			panic(err)
		}
		c.Close()
	}
}

func BenchmarkXGB_Python_UNIXsocket_RawBytes_NewConnection(b *testing.B) {
	// start server
	cmd := exec.Command("python3", "../python-server/xgb-unixsocket-rawbytes-newconn.py", "../sc", "../data/models/13features.xgb")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer func() { os.Remove("../sc") }() // remove socket

	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err != nil {
		panic(err)
	}
	defer syscall.Kill(-pgid, 15) // kill server

	time.Sleep(1 * time.Second)
	benchmarkUNIXSocketRawBytesNewConn(b, "../sc")
}
