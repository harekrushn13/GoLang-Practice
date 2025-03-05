package main

import (
	"fmt"
	"reflect"
)

// Define a struct with different tag values
type Product struct {
	ID    int    `alias:"product_id"`
	Name  string `alias:""`
	Code  string `alias:"product_code"`
	Price float64
}

func main() {
	// Create an instance of the Product struct
	p := Product{
		ID:    123,
		Name:  "Laptop",
		Code:  "LP123",
		Price: 123.456,
	}

	// Get the type of the struct
	t := reflect.TypeOf(p)

	// Iterate over the fields and inspect the tags
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Lookup tag values
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println("Alias:", alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}
}
