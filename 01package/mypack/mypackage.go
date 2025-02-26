package mypack

import "fmt"

// Global variable
var message = "Initializing mypackage"

func init() {
	fmt.Println("mypackage init() called")
}

func Hello() {
	fmt.Println("Hello from mypackage!")
}
