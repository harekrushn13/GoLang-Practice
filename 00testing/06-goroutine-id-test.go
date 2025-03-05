package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getGoroutineID() int {
	buf := make([]byte, 64)               // Create buffer
	buf = buf[:runtime.Stack(buf, false)] // Get stack trace
	stack := string(buf)
	fmt.Println("stack:", stack)
	// Extract Goroutine ID from "goroutine <ID> [running]:"
	parts := strings.Split(stack, " ")
	id, _ := strconv.Atoi(parts[1])
	return id
}

func main() {
	// 1
	fmt.Println("Main Goroutine ID:", getGoroutineID())

	go func() {
		fmt.Println("New Goroutine ID:", getGoroutineID())
	}()

	// Allow time for goroutine to print
	runtime.Gosched()

	// 2
	//go leakyGoroutine()
	//go leakyGoroutine()
	//
	//fmt.Println("Server started on :6060")
	//http.ListenAndServe(":6060", nil) // Expose pprof
}

func leakyGoroutine() {
	for {
		time.Sleep(1 * time.Second)
	}
}
