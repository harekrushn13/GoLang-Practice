package main

/*
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

// Function to get the current thread id (for Linux)
unsigned long getThreadID() {
	return (unsigned long)pthread_self();
}
*/
import "C"
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	threadID := C.getThreadID()
	fmt.Printf("Locked to thread. Goroutine ID: %s, OS Thread ID: %d\n", getGoroutineID(), threadID)

	runtime.Gosched()

	threadID = C.getThreadID()
	fmt.Printf("Goroutine resumed. OS Thread ID: %d\n", threadID)

	time.Sleep(2 * time.Second)

	threadID = C.getThreadID()
	fmt.Printf("Main goroutine completed. OS Thread ID: %d\n", threadID)
}

func getGoroutineID() string {
	return "Simulated Goroutine ID"
}
