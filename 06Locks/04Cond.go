package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var ready = false

func waitForCondition(wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Condition met, proceeding...")
	cond.L.Unlock()
}

func signalCondition() {
	cond.L.Lock()
	ready = true
	cond.L.Unlock()
	cond.Broadcast()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go waitForCondition(&wg)
	go waitForCondition(&wg)

	time.Sleep(2 * time.Second)
	signalCondition()

	wg.Wait()
}
