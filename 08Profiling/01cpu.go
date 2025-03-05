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

	for i := 0; i < 100; i++ {
		//doWork(i)
		heavyComputation()
		heavyComputation2()
	}
	//heavyComputation()
	//time.Sleep(time.Second)
}

func heavyComputation() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i * i
	}
	//fmt.Println("Computation done:", sum)
}
func heavyComputation2() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i * i * 1
	}
	//fmt.Println("Computation done:", sum)
}
