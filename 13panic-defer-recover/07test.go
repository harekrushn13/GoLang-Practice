package main

import (
	"fmt"
	"runtime/debug"
	"sync"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
		fmt.Println(string(debug.Stack()))
	}()

	var wg sync.WaitGroup // Declare a WaitGroup

	wg.Add(1) // Add 1 before calling the goroutine

	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println(r)
			}
			fmt.Println(string(debug.Stack()))
		}()
		wg.Done()
		wg.Wait()
		wg.Done()
	}()

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Main function execution completed!")
}
