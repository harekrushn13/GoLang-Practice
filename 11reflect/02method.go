package main

import (
	"fmt"
	"reflect"
)

// Define a struct
type MyStruct struct{}

// Define methods for MyStruct
func (s MyStruct) ExportedMethod() {
	fmt.Println("This is an exported method.")
}

func (s MyStruct) unExportedMethod() {
	fmt.Println("This is an unexported method.")
}

func main() {
	// Create an instance of MyStruct
	instance := MyStruct{}

	// Get the type of MyStruct
	t := reflect.TypeOf(instance)
	fmt.Println(t.Name(), t.NumMethod())

	// Iterate over all methods of MyStruct
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)

		// Check if the method is exported
		fmt.Printf("Method: %s, IsExported: %v\n", method.Name, method.IsExported())
	}

}
