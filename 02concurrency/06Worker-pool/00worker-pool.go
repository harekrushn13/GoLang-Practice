package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numWorkers := 3
	numTasks := 10

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	for i := 0; i < numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()

	close(results)

	for v := range results {
		fmt.Printf("Result : %d\n", v)
	}
}

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range tasks {
		fmt.Printf("worker %d received task #%d\n", id, v)
		results <- v * v

		time.Sleep(1 * time.Second)
	}
}
