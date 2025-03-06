package main

import "fmt"

type MyInt = int
type MyIntAlias = MyInt

func main() {
	var a MyInt = 5
	var b MyIntAlias = a // Alias type can be assigned directly

	fmt.Printf("%T %T\n", a, b)
	fmt.Println(a, b)
}
