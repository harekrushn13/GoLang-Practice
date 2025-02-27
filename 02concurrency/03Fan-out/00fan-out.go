package main

import (
	"fmt"
)

func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	input := generateWork(work)

	out1 := fanOut(input)
	out2 := fanOut(input)
	out3 := fanOut(input)
	out4 := fanOut(input)

	for range work {
		select {
		case value := <-out1:
			fmt.Println("Output 1 got:", value)
		case value := <-out2:
			fmt.Println("Output 2 got:", value)
		case value := <-out3:
			fmt.Println("Output 3 got:", value)
		case value := <-out4:
			fmt.Println("Output 4 got:", value)
		}
	}
}

func fanOut(input <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		//defer close(out)

		for data := range input {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		//defer close(ch)

		for _, v := range work {
			ch <- v
		}
	}()

	return ch
}
