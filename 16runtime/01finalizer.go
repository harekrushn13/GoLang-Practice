package main

import (
	"fmt"
	"runtime"
	"time"
)

type Resource struct {
	ID int
}

func (r *Resource) cleanup() {
	fmt.Printf("Cleaning up resource with ID %d\n", r.ID)
}

func main() {
	// Create an object of Resource
	resource := &Resource{ID: 1}

	// Set the finalizer for the resource object
	// When the resource is garbage collected, the cleanup method will be called.
	runtime.SetFinalizer(resource, func(r *Resource) {
		fmt.Println("jdis")
		r.cleanup()
	})

	// Simulate the resource being used
	fmt.Println("Using resource with ID:", resource.ID)

	// Explicitly set resource to nil to allow GC (garbage collection) to collect it
	resource = nil

	// Suggest to the garbage collector to run, but we can't guarantee it will happen immediately
	// This is only to demonstrate the finalizer call
	runtime.GC()

	// Give the garbage collector some time to collect
	// It's not guaranteed that the finalizer will run immediately.
	time.Sleep(50 * time.Second)
}
