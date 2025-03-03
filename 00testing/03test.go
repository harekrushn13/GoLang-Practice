package main

import (
	"fmt"
	"os"
)

func main() {
	// Open file in read-only mode
	file, err := os.OpenFile("./00testing/02file.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Try to write to the file
	_, _ = file.Write([]byte("Hello, world!"))
	//if err != nil {
	//	// This will print an error because the file is open in read-only mode
	//	fmt.Println("sjaijsi- jjisja")
	//	fmt.Println("Error writing to file:", err)
	//	fmt.Println("sjaijsi- jjisja-qjijqiw")
	//}

	fmt.Println("sjaijsi")
}
