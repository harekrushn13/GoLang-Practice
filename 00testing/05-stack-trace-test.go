package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// stack trace

func main() {
	causePanic()
}

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack Trace:")
			fmt.Println(string(debug.Stack())) // Print stack trace
			fmt.Println("-------------------------")

		}
	}()

	triggerError()
	//captureStackTrace()
}

func triggerError() {
	panic("Something went wrong!")
}

func captureStackTrace() {
	pc := make([]uintptr, 10) // Store up to 10 stack frames
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		fmt.Printf("Function: %s\nFile: %s:%d\n\n", frame.Function, frame.File, frame.Line)
	}
}
