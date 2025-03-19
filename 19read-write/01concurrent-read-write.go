package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := "./19read-write/data.txt"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	os.WriteFile(filePath, []byte("Hello World\n"), 0777)

	for i := 0; i < 10; i++ {
		go writeToFile(file, fmt.Sprintf("Writer %d\n", i+1))
		go readFromFile(file)
	}

	time.Sleep(10 * time.Second)
}

func writeToFile(file *os.File, content string) {

	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Error while writing to file:", err)
	} else {
		fmt.Println("Write:", content)
	}
}

func readFromFile(file *os.File) {
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	fmt.Println("Read:", string(buffer[:n]))
}
