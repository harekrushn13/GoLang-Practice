package main

import (
	"fmt"
	"time"
)

func main() {
	// 1
	//You can create a time.Timer using time.NewTimer(duration), which waits for the duration and then sends a signal to its channel.
	//timer := time.NewTimer(3 * time.Second)
	//
	//fmt.Println("Waiting for timer...")
	//<-timer.C // Blocks until the timer expires
	//fmt.Println("Timer expired at", time.Now())

	// 2
	// If you don’t need to stop or reset the timer, use time.After(duration), which returns a channel that receives a value after the duration.
	// This is simpler than time.NewTimer when you just need a one-time delay.
	//fmt.Println("Waiting for 2 seconds...")
	//<-time.After(2 * time.Second)
	//fmt.Println("Done waiting at", time.Now())

	// 3
	// If you create a time.Timer but no longer need it, you must stop it to prevent memory leaks.
	//timer := time.NewTimer(5 * time.Second)
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	stopped := timer.Stop()
	//	if stopped {
	//		fmt.Println("Timer stopped before it expired.")
	//	}
	//}()
	//
	//fmt.Println("Waiting for timer...")
	//<-timer.C // Will not execute if stopped
	//fmt.Println("Timer expired (should not print)")

	// 4
	// Instead of creating a new timer every time, you can reuse an existing one using Reset().
	//timer := time.NewTimer(3 * time.Second)
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	timer.Reset(4 * time.Second) // Resets timer before it expires
	//	fmt.Println("Timer reset!")
	//}()
	//
	//<-timer.C
	//fmt.Println("Timer expired at", time.Now()) // Fires after 4 seconds, not 3

	// Edge case : Reset() does not work on an expired timer unless you first drain its channel.
	//timer := time.NewTimer(1 * time.Second)
	//time.Sleep(2 * time.Second) // Timer already expired
	//
	//timer.Reset(3 * time.Second) // Resetting a dead timer
	//x := <-timer.C
	//fmt.Println("Timer expired (unexpected behavior)", x)

	// Solution: Drain the Channel First
	//timer := time.NewTimer(1 * time.Second)
	//<-timer.C // Drain the expired timer event
	//
	//timer.Reset(3 * time.Second) // Now Reset works properly
	//<-timer.C
	//fmt.Println("Timer expired correctly")

	// 5 : Using Timer with select
	// time.Timer is often used inside a select statement to time out operations.

	timer := time.NewTimer(3 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(4 * time.Second) // Simulate work
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task finished in time!")
	case <-timer.C:
		fmt.Println("Timeout! Task took too long.")
	}

}
