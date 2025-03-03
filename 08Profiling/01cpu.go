package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func doWork(i int) {
	_ = i * i * i * i * i
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1e9; i++ {
		doWork(i)
	}
}
