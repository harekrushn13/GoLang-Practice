package main

import (
	"fmt"
	"time"
)

func printNumbers() {

	x := 1
	for i := 1; i <= 3; i++ {
		defer func() {
			fmt.Println("Deferred:", i, x)
		}()
		x++

		go func() {
			fmt.Println("Goroutine:", i, x)
		}()
	}
}

//func printNumbers() {
//	for i := 1; i <= 3; i++ {
//		defer fmt.Println("Deferred:", i)
//		go func(n int) {
//			fmt.Println("Goroutine:", n)
//		}(i) // Pass `i` as an argument
//	}
//}

func main() {
	printNumbers()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed")
}
