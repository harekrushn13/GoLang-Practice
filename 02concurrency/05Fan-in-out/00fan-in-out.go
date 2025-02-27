package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := generateWork([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	sq1 := square(numbers)
	sq2 := square(numbers)

	merged := merge(sq1, sq2)

	for v := range merged {
		fmt.Println(v)
	}
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, channel := range channels {
		wg.Add(1)

		go func(ch <-chan int) {
			defer wg.Done()

			for v := range ch {
				out <- v
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func square(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range ch {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func generateWork(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	return out
}
