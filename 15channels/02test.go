package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Println("start")
	ch <- 1
	fmt.Println("end")
}
