package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1)

	slice := make([]int, 1000)

	var wg sync.WaitGroup

	wg.Add(len(slice))
	for i := 0; i < len(slice); i++ {
		go func() {
			defer wg.Done()
			slice[0] = slice[0] + 1
			fmt.Printf("Written %d to index %d\n", slice[i], i)
		}()
	}

	wg.Add(len(slice))
	for i := 0; i < len(slice); i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("Read %d from index %d\n", slice[i], i)
		}()
	}

	wg.Wait()

	fmt.Println(slice)
	fmt.Println("Completed without fatal error.")
}
