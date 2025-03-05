package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	printGoEnv()

	fmt.Println("Changing GOPATH at runtime...")
	os.Setenv("GOPATH", "/new/gopath")
	os.Setenv("GOVERSION", "go1.1")

	printGoEnv()

}

func printGoEnv() {
	fmt.Println("=== Current Go Environment ===")
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("GOVERSION: %s\n", runtime.Version())
	fmt.Printf("GOPATH: %s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("GOVERSION: %s\n", os.Getenv("GOVERSION"))

	time.Sleep(5 * time.Second)
	fmt.Println("==============================")

}
