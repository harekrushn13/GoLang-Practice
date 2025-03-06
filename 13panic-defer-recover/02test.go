package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("First recover:", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Second recover:", r)
		}
	}()

	panic("Error occurred")
}
