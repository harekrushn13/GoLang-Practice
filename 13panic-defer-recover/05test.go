package main

import "fmt"

// Function to demonstrate recovery from panic
func safeFunction() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from safeFunction:", r)
		}
	}()

	fmt.Println("Executing safeFunction...")
	panic("Something went wrong in safeFunction1!") // Triggers panic
	panic("Something went wrong in safeFunction2!") // Triggers panic
}

// Function to demonstrate safe division
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from division error:", r)
		}
	}()
	//defer func() {
	//	recover()
	//}()
	//defer fmt.Println("Recovered:", recover())
	fmt.Println("Result:", a/b) // If b = 0, panic occurs
}

func triggerPanic() {
	defer fmt.Println("triggerPanic statement executed!")
	triggerPanicinside()
	fmt.Println("This will not be printed due to panic")
}

func triggerPanicinside() {
	defer fmt.Println("triggerPanicinside statement executed!")
	panic("Something went wrong33455!")
	fmt.Println("This will not be printed due to panic")
}

func processTask() {
	defer fmt.Println("Defer statement executed! Cleanup done!")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	triggerPanic()

	fmt.Println("This will not be printed due to panic")
}

func main() {
	fmt.Println("Start of main function")
	// Recover from panic inside a function
	safeFunction()

	// Safe division example
	safeDivide(10, 2)
	safeDivide(10, 0) // This will be caught by recover

	// Normal execution continues
	fmt.Println("Program continues running after recover handling")

	// ------------------------------ //

	// Start processing the task
	fmt.Println("Starting task processing...")
	processTask()
	fmt.Println("After recovery, the program continues execution here?")
	fmt.Println("Task completed with recovery.")
}
