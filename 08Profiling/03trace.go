package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	// Simulate workload
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
}
