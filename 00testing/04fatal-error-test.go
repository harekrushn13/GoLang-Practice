package main

import (
	"fmt"
	"sync"
)

func infiniteRecursion() {
	fmt.Println("Recursing...")
	infiniteRecursion() // No base case, causes stack overflow
}

// fatal errors : are not recoverable

func main() {
	// stack ovreflow
	//infiniteRecursion()

	//out of memory
	//bigSlice := make([]int, 1<<60) // Allocating an extremely large slice
	//_ = bigSlice

	//cocurrentmap writes
	//m := make(map[int]int)
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		m[i] = i * 10 // Multiple Goroutines writing to the same map
	//	}(i)
	//}
	//
	//wg.Wait()
	//fmt.Println(m)

	// corrupt memory
	//runtime.GC() // Explicit garbage collection (can cause issues in unsafe memory use)
	//fmt.Println("Program running...")

	// deadlocks
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ch <- 42 // No Goroutine to receive data, causes deadlock
		wg.Done()
	}()

	wg.Wait() // Main Goroutine is blocked forever
}
