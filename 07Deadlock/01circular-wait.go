package main

import (
	"fmt"
	"sync"
	"time"
)

var mu1, mu2 sync.Mutex

func task1() {
	mu1.Lock()
	time.Sleep(1 * time.Second) // Simulating work
	mu2.Lock()                  // Deadlock if task2 already holds mu2
	fmt.Println("Task 1 executed")
	mu2.Unlock()
	mu1.Unlock()
}

func task2() {
	mu2.Lock()
	time.Sleep(1 * time.Second) // Simulating work
	mu1.Lock()                  // Deadlock if task1 already holds mu1
	fmt.Println("Task 2 executed")
	mu1.Unlock()
	mu2.Unlock()
}

func main() {
	go task1()
	go task2()
	time.Sleep(3 * time.Second)
}
