package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
		return
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		fmt.Println("Error starting trace:", err)
		return
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
