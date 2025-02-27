package main

import (
	"fmt"
	"sync"
)

const (
	size      = 1000
	workerNum = 10
	chunkSize = 100
)

func main() {
	input := generator()

	result := make(chan float64)

	var wg sync.WaitGroup

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go worker(input, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	totalSum := 0.0
	for sum := range result {
		totalSum += sum
	}

	average := totalSum / (size / chunkSize)

	fmt.Printf("Average : %.2f\n", average)
}

func worker(ch <-chan []int, result chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for x := range ch {

		for _, y := range x {

			sum += y
		}
	}
	result <- float64(sum) / chunkSize
}

func generator() <-chan []int {
	ch := make(chan []int)

	data := make([]int, size)

	for i := 0; i < size; i++ {

		data[i] = i + 1
	}

	go func() {

		for i := 0; i < size/chunkSize; i++ {

			start := i * chunkSize

			end := start + chunkSize

			ch <- data[start:end]
		}

		close(ch)
	}()

	return ch
}
