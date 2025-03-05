package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var mu sync.Mutex

func work(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	time.Sleep(100 * time.Millisecond)
	mu.Unlock()
}

func main() {
	runtime.SetBlockProfileRate(1)

	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Println("Error creating block profile file:", err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg)
	}

	wg.Wait()

	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Println("Error writing block profile:", err)
	}
}
