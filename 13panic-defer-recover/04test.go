package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("This will execute last1")
	defer fmt.Println("This will execute last2")
	defer fmt.Println("This will execute last3")

	fmt.Println("This will execute first")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close() // Ensures file is closed when function exits

	fmt.Println("File opened successfully")

}
