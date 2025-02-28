package main

import "fmt"

func main() {
	s := make([]float64, 100)
	fmt.Println(len(s), cap(s))

	func(s []float64) {
		fmt.Println(len(s), cap(s))
		s = append(s, 101.9)
		fmt.Println(len(s), cap(s))
	}(s)

	fmt.Println(len(s), cap(s))

}
