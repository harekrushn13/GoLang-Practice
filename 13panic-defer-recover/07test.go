package main

import (
	"fmt"
	"sync"
)

func main() {
	//defer func() {
	//	r := recover()
	//	if r != nil {
	//		fmt.Println("Main recover : ", r)
	//	}
	//	fmt.Println("Main Stacktrace : ", string(debug.Stack()))
	//}()

	var wg sync.WaitGroup // Declare a WaitGroup

	wg.Add(1) // Add 1 before calling the goroutine

	go func(wg *sync.WaitGroup) {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("Routine recover : ", r)
			}
			//fmt.Println("Routine Stacktrace : ", string(debug.Stack()))
		}()
		wg.Done()
		wg.Wait()
		wg.Done()
	}(&wg)

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Main function execution completed!")
}
