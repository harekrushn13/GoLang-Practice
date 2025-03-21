package main

import (
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
		fmt.Println("Checkpoint 1")
		panic(1)
	}()
	defer func() {
		r := recover()
		fmt.Println(r)
		fmt.Println("Checkpoint 2")
		panic(2)
	}()
	panic(999)
}
