package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 1000000)
	_ = data

	pprof.WriteHeapProfile(f)
}
