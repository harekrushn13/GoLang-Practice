package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 2)

	go producer(queue)

	//time.Sleep(time.Second)
	consumer(queue)
}

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Producing task : %d\n", i)
		queue <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(queue)
}

func consumer(queue <-chan int) {
	for task := range queue {
		fmt.Printf("Consuming task : %d\n", task)
		time.Sleep(1 * time.Second)
	}
}
