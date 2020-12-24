package fifo_test

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

const fifoFeaturesPath = "features.fifo"
const fifoPredictionsPath = "predictions.fifo"

func predict(features []float32) (float32, error) {
	fifoFeatures, err := os.OpenFile(fifoFeaturesPath, os.O_WRONLY, 0666)
	if err != nil {
		return 0, fmt.Errorf("can not open fifo features for write: %w", err)
	}

	if err := binary.Write(fifoFeatures, binary.LittleEndian, features); err != nil {
		return 0, fmt.Errorf("can not write features to fifo: %w", err)
	}

	fifoFeatures.Close()

	fifoPredictionns, err := os.OpenFile(fifoPredictionsPath, os.O_RDONLY, 0666)
	if err != nil {
		return 0, fmt.Errorf("can not open fifo predictions for read: %w", err)
	}

	var prediction float32
	if err := binary.Read(fifoPredictionns, binary.LittleEndian, &prediction); err != nil {
		return 0, fmt.Errorf("can not read predictions from fifo: %w", err)
	}

	fifoPredictionns.Close()

	return prediction, nil
}

func BenchmarkFifoXGBoostC(b *testing.B) {
	syscall.Mkfifo(fifoFeaturesPath, 0666)
	syscall.Mkfifo(fifoPredictionsPath, 0666)

	features := []float32{72.0, 72.0, 69.0, 60.0, 7.0, 0, 0, 0, 1, 1, 0, 1, 0}

	// Start XGBoost C server listening to UNIX pipes
	cmd := exec.Command("make", "run", "-C", "xgboost")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Kill XGBoost C server once finish
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err != nil {
		panic(err)
	}
	defer syscall.Kill(-pgid, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err := predict(features)
		if result == 0 || err != nil {
			panic(err)
		}
	}
}
