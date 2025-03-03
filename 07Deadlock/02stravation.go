package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func starvedWorker(id int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Worker", id, "starting")
	time.Sleep(2 * time.Second) // Long execution time
	fmt.Println("Worker", id, "finished")
}

func main() {
	for i := 0; i < 3; i++ {
		go starvedWorker(i) // Only first goroutine gets executed quickly
	}
	time.Sleep(5 * time.Second)
}
