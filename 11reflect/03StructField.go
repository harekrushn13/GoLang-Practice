package main

import (
	"fmt"
	"reflect"
)

// Define a struct with fields and tags
type Person struct {
	Name    string `json:"name" xml:"name"`
	Age     int    `json:"age"`
	address string `json:"address,omitempty"`
}

func main() {
	// Create an instance of the struct
	p := Person{
		Name:    "John Doe",
		Age:     30,
		address: "1234 Main St",
	}

	// Get the type of the struct
	t := reflect.TypeOf(p)

	// Iterate over the fields in the struct
	for i := 0; i < t.NumField(); i++ {
		// Get the field information
		field := t.Field(i)

		// Print field name, type, visibility, and tags
		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", field.Type)
		fmt.Printf("Is Exported: %v\n", field.IsExported())
		fmt.Printf("Tag (json): %s\n", field.Tag.Get("json"))
		fmt.Printf("Tag (xml): %s\n", field.Tag.Get("xml"))
		fmt.Println()
	}

	// Using VisibleFields function to check all visible fields, including unexported ones
	fmt.Println("Visible Fields:")
	tVisibleFields := reflect.VisibleFields(t)
	for _, f := range tVisibleFields {
		fmt.Printf("Visible Field: %s, Is Exported: %v\n", f.Name, f.IsExported())
	}
}
