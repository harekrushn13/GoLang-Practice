package main

import "fmt"

func panicInsidemutipleDefer() {
	defer func() {
		fmt.Println("\n--- Panic Inside1 Defer ---")
		recover()
	}()
	defer func() {
		fmt.Println("\n--- Panic Inside2 Defer ---")
		panic("Panic4 inside defer!")
		panic("Panic3 inside defer!")
	}()
	defer func() {
		fmt.Println("\n--- Panic Inside3 Defer ---")
		panic("Panic2 inside defer!")
		panic("Panic1 inside defer!")
	}()
	panic("Original panic in panicInsideDefer()")
}

func main() {
	panicInsidemutipleDefer()

	fmt.Println("continue")
}
