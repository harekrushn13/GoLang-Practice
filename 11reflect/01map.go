package main

import (
	"fmt"
	"reflect"
)

func main() {
	map1 := map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}

	val1 := reflect.ValueOf(map1)
	fmt.Println("reflect.ValueOf(map1) : ", val1)

	iter := val1.MapRange()
	fmt.Println("val1.MapRange() : ", iter)

	fmt.Println("Iterating over map1:")
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Printf("Key: %v, Value: %v\n", key.Interface(), value.Interface())
	}

	map2 := map[string]int{
		"dog":   20,
		"cat":   30,
		"mouse": 40,
	}

	val2 := reflect.ValueOf(map2)
	iter.Reset(val2)

	fmt.Println("\nIterating over map2 after Reset:")
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Printf("Key: %v, Value: %v\n", key.Interface(), value.Interface())
	}

	if !iter.Next() {
		fmt.Println("\nIterator is exhausted:", iter.Next())
	}
}
