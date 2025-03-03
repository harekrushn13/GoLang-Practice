package main

import (
	"fmt"
	"sync"
)

var (
	mu  sync.Mutex
	cnt int
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	cnt++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", cnt)
}
