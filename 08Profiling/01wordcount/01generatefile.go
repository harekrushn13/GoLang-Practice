package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	fileName     = "./01wordcount/sample.txt" // Output file name
	fileSizeMB   = 50                         // Desired file size in MB
	chunkSize    = 1024 * 1024                // 1MB chunks
	wordListSize = 1000                       // Number of different words
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a list of random words
	words := generateRandomWords(wordListSize)

	// Create and open the file
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the file in chunks
	totalSize := 0
	for totalSize < fileSizeMB*1024*1024 {
		chunk := generateChunk(words, chunkSize)
		n, err := file.WriteString(chunk)
		if err != nil {
			panic(err)
		}
		totalSize += n
	}
}

// generateRandomWords creates a list of `count` random words.
func generateRandomWords(count int) []string {
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = randomWord()
	}
	return words
}

// randomWord generates a random word of length 3 to 8.
func randomWord() string {
	length := rand.Intn(6) + 3 // Word length between 3 and 8
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(byte(rand.Intn(26) + 'a'))
	}
	return sb.String()
}

// generateChunk creates a chunk of random words up to `size` bytes.
func generateChunk(words []string, size int) string {
	var sb strings.Builder
	sb.Grow(size)

	for sb.Len() < size-10 { // Leave space for the last word
		sb.WriteString(words[rand.Intn(len(words))])
		sb.WriteByte(' ') // Space between words
	}

	return sb.String()
}
