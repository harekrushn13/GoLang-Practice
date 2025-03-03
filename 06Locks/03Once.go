package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initOnce() {
	fmt.Println("Initialization done.")
}

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ", i)
	once.Do(initOnce)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
