package main

import (
	"01package/mypack"
	"fmt"
	"runtime"
)

var globalVar = "Global variable in main"

func init() {
	fmt.Println("main package init() called")
}

func main() {
	fmt.Println("main() function started")
	mypack.Hello()
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Heap Allocation: %v KB\n", memStats.HeapAlloc/1024)
	runtime.GC()

	//fmt.Println("Max float32:", math.MaxFloat32)
	//fmt.Println("Max float64:", math.MaxFloat64)
}
