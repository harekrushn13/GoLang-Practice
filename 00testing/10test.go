package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Process ID:", file.Name())
		}
	}
}
