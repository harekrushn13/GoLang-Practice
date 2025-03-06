package main

import "fmt"

func save() {
	recover()
}

func f1() {
	defer save()
	defer fmt.Println("f1: Defer 1")
	defer fmt.Println("f1: Defer 2")
	panic("Panic in f1")
}

func f2() {
	defer fmt.Println("f2: Defer 1")
	f1()
	defer fmt.Println("f2: Defer 2")
}

func main() {
	defer fmt.Println("main: Defer 1")
	f2()
	defer fmt.Println("main: Defer 2")
}
