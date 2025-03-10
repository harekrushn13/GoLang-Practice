package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// 1
	//for i := 0; i < 5; i++ {
	//	<-ticker.C // ticker.C is a channel that receives values at the specified interval.
	//	fmt.Println("Tick at:", time.Now())
	//}

	// 2
	// If a Ticker event is blocked (e.g., due to a slow function), ticks can pile up.
	//The program misses ticks since time.Sleep(3 * time.Second) takes longer than the ticker interval.
	//for v := range ticker.C {
	//	fmt.Println("Tick ", v)
	//	time.Sleep(3 * time.Second) // Blocking the ticker
	//}

	// 3 : non-blocking solution
	//go func() {
	//	for range ticker.C {
	//		go func() {
	//			fmt.Println("Tick")
	//		}()
	//	}
	//}()
	//
	//time.Sleep(5 * time.Second)

	// 4 :Ticker with select
	//	The select statement listens to both ticker.C and done channels.
	//	The ticker stops when the done signal is received.
	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick at", time.Now())
		case <-done:
			fmt.Println("Done!")
			return
		}
	}
}
