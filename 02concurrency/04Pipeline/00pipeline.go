package main

import "fmt"

func main() {
	input := generateWork([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	output := filter(input)
	//output = square(input)
	//output = half(input)

	for v := range output {
		fmt.Println(v)
	}
}

func filter(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range input {
			if n%2 == 0 {
				out <- n
			}
		}
	}()

	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range input {
			out <- n * n
		}
	}()

	return out
}

func half(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range input {
			out <- n / 2
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for _, x := range work {
			c <- x
		}
	}()

	return c
}
