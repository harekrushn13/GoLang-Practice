package main

import (
	"fmt"
	"sync"
)

var cmap sync.Map

func main() {
	cmap.Store("key1", 42)
	cmap.Store(12, "ndiw")

	val, ok := cmap.Load("key1")
	if ok {
		fmt.Println("Loaded value:", val)
	}

	cmap.Delete("key1")

	fmt.Println(cmap.Load(12))
}
