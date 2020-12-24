package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {
	c, err := net.Dial("unix", "../xgb-go.socket")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	for i := 0; i < 1; i++ {
		features := []float64{13}
		if err := binary.Write(c, binary.LittleEndian, features); err != nil {
			log.Fatal(err)
		}

		predictions := make([]float32, 1)
		if err := binary.Read(c, binary.LittleEndian, predictions); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("predictions %#v\n", predictions)
	}
}
