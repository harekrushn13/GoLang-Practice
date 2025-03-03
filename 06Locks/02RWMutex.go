package main

import (
	"fmt"
	"sync"
)

var (
	rwMu sync.RWMutex
	data = make(map[string]int)
)

func readData(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.RLock()
	fmt.Println("Reading data : ", data)
	rwMu.RUnlock()
}

func writeData(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMu.Lock()
	data["x"]++
	rwMu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go writeData(&wg)
		go readData(&wg)
	}

	wg.Wait()
	fmt.Println(data)
}
