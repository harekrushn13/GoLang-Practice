package main

import (
	"fmt"
	"time"
)

func normalPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in normalPanic():", r)
		}
	}()
	fmt.Println("\n--- Normal Panic Example ---")
	panic("Panic in normalPanic() function!")
	fmt.Println("This will not execute")
}

func panicInsideDefer() {
	defer func() {
		fmt.Println("\n--- Panic Inside1 Defer ---")
		panic("Panic inside defer!") // Overwrites previous panic
	}()
	panic("Original panic in panicInsideDefer()") // This gets replaced
}

func multipleDefers() {
	fmt.Println("\n--- Multiple Defers Example ---")
	defer fmt.Println("Deferred function 1")
	defer fmt.Println("Deferred function 2")
	defer fmt.Println("Deferred function 3")
	panic("Panic in multipleDefers()") // Triggers deferred calls
	defer fmt.Println("Deferred function 4")
	defer fmt.Println("Deferred function 5")
}

func recoverOnlyOnce() {
	defer func() {
		if r := recover(); r == nil {
			fmt.Println("\n--- Recover already done ---")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n--- Recover Example ---")
			fmt.Println("Recovered:", r)
		}
	}()

	panic("Recover example panic!")
}

func panicInGoroutine() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("\n--- Recovered Inside Goroutine ---")
				fmt.Println("Recovered:", r)
			}
		}()
		fmt.Println("Worker Goroutine Started")
		panic("Panic inside Goroutine!")
	}()

	time.Sleep(time.Second) // Wait for goroutine execution
}

func unhandledGoroutinePanic() {
	go func() {
		fmt.Println("\n--- Unhandled Goroutine Panic ---")
		panic("This panic will crash the program if unhandled!")
	}()
	time.Sleep(time.Second) // Give time for panic to execute
}

func main() {
	fmt.Println("=== Go Defer, Panic & Recover Examples ===")

	// Example 1: Panic with Recovery
	//normalPanic()

	// Example 2: Panic inside a deferred function
	// Uncommenting this will crash the program
	//panicInsideDefer()

	// Example 3: Multiple defer statements executing in LIFO order
	//multipleDefers()

	// Example 4: Only one recover() executes
	//recoverOnlyOnce()

	// Example 5: Panic inside a Goroutine (Recovered)
	//panicInGoroutine()

	// Example 6: Unhandled Goroutine Panic (Crashes program)
	// Uncommenting this will crash the program
	//unhandledGoroutinePanic()

	fmt.Println("\nMain function completed successfully!")
}
