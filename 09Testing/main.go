package main

import (
	"fmt"
	"mytesting/mathpack"
)

func main() {
	result := mathpack.Add(2, 3)
	fmt.Println(result)

	res, _ := mathpack.FetchUser(1)
	fmt.Println(res)
}
