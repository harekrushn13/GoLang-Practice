package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"unicode"
)

func main() {
	// Create CPU profiling file
	cpuProfile, err := os.Create("./01wordcount/cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	pprof.StartCPUProfile(cpuProfile) // Start CPU profiling
	defer pprof.StopCPUProfile()      // Stop profiling when function exits

	// Open the file
	file, err := os.Open("./01wordcount/sample.txt") // Change filename as needed
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Count words in file
	totalWords, err := countTotalWords(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print total word count
	fmt.Println("Total Words:", totalWords)
}

// countTotalWords reads a file and counts total words using only `os` package.
func countTotalWords(file *os.File) (int, error) {
	// Read entire file into memory
	data := make([]byte, 1024) // Read in 1KB chunks
	totalWords := 0
	inWord := false

	for {
		n, err := file.Read(data)
		if n == 0 {
			break // End of file
		}
		if err != nil {
			return 0, err
		}

		// Process each byte in the chunk
		for i := 0; i < n; i++ {
			if unicode.IsSpace(rune(data[i])) {
				if inWord {
					totalWords++
					inWord = false
				}
			} else {
				inWord = true
			}
		}
	}

	// Count the last word if the file doesn't end in whitespace
	if inWord {
		totalWords++
	}

	return totalWords, nil
}
