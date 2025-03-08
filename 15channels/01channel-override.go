package main

import (
	"fmt"
	"time"
)

// G1
func main() {
	tasks := []string{"task1", "task2", "task3"}

	ch := make(chan string, 3)
	fmt.Println(ch, &ch)

	for i, task := range tasks {
		fmt.Println(&tasks[i], &task, &i)
		//go func() {
		//	time.Sleep(1 * time.Second)
		//	fmt.Println(task, &task)
		//}()
		ch <- task
	}

	go worker(&ch)
	time.Sleep(2 * time.Second)
}

// G2
func worker(ch *chan string) {
	fmt.Println(ch, &ch)

	// ch = make(chan string, 3)
	*ch = make(chan string)

	fmt.Println(ch, &ch)
	for {
		t := <-*ch
		process(t)
	}
}

func process(t string) {
	fmt.Println(t)
}
