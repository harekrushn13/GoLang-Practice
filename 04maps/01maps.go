package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create a map for demonstration
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// 1. Clear
	//fmt.Println("--- Clear ---")
	//fmt.Println("Original Map:", m)
	//maps.Clear(m) // Removes all entries from the map
	//fmt.Println("After Clear:", m) // Output: map[] (empty map)

	// Edge Case: Clear on nil map
	// var nilMap map[string]int
	// maps.Clear(nilMap) // This will panic at runtime

	// 2. Clone
	fmt.Println("\n--- Clone ---")
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	cloned := maps.Clone(m) // Creates a shallow copy of the map
	fmt.Println("Original Map:", m)
	fmt.Println("Cloned Map:", cloned)

	cloned["a"] = 11
	fmt.Println("After cloned[\"a\"] = 11 Original Map:", m)
	fmt.Println("After cloned[\"a\"] = 11 Cloned Map:", cloned)

	// Edge Case: Clone on nil map
	var nilMap map[string]int
	clonedNil := maps.Clone(nilMap)
	fmt.Println("Cloned Nil Map:", clonedNil) // Output: map[] (nil map)

	mpp := map[string]map[string]int{
		"M1": map[string]int{
			"Key1": 10,
			"Key2": 100,
		},
		"M2": map[string]int{
			"Key1": 100,
			"Key2": 1000,
		},
	}

	//cmpp := maps.Clone(mpp)
	cmpp := mpp
	cmpp["M1"]["Key1"] = 10000
	cmpp["M2"] = map[string]int{
		"New": 1,
	}
	fmt.Println("mpp: ", mpp)
	fmt.Println("cmpp: ", cmpp)

	// 3. Copy
	fmt.Println("\n--- Copy ---")
	dst := map[string]int{"x": 10, "y": 20}
	src := map[string]int{"y": 200, "z": 300}
	fmt.Println("Destination Map (Before Copy):", dst)
	fmt.Println("Source Map:", src)
	maps.Copy(dst, src)                               // Copies all key/value pairs from src to dst
	fmt.Println("Destination Map (After Copy):", dst) // Output: map[x:10 y:200 z:300]

	// Edge Case: Copy with nil maps
	// var nilDst map[string]int
	// maps.Copy(nilDst, src) // This will panic at runtime

	// 4. DeleteFunc
	fmt.Println("\n--- DeleteFunc ---")
	m = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println("Original Map:", m)
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 == 0 // Delete keys with even values
	})
	fmt.Println("After DeleteFunc:", m) // Output: map[a:1 c:3]

	// Edge Case: DeleteFunc with nil map or nil function
	// var nilMap2 map[string]int
	// maps.DeleteFunc(nilMap2, func(k string, v int) bool { return true }) // Panics
	// maps.DeleteFunc(m, nil) // Panics

	// 5. Equal
	fmt.Println("\n--- Equal ---")
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}
	fmt.Println("m1 == m2:", maps.Equal(m1, m2)) // Output: true
	fmt.Println("m1 == m3:", maps.Equal(m1, m3)) // Output: false

	// Edge Case: Equal with nil maps
	//var nilMap3 map[string]int
	//fmt.Println("nilMap3 == nil:", maps.Equal(nilMap3, nil)) // Output: true

	// 6. EqualFunc
	fmt.Println("\n--- EqualFunc ---")
	m4 := map[string]int{"a": 1, "b": 2}
	m5 := map[string]int{"a": 1, "b": 20}
	equal := maps.EqualFunc(m4, m5, func(v1, v2 int) bool {
		return v1%2 == v2%2 // Compare values based on parity (even/odd)
	})
	fmt.Println("m4 and m5 have values with same parity:", equal) // Output: true

	// Edge Case: EqualFunc with nil maps or nil function
	//var nilMap4 map[string]int
	//fmt.Println("nilMap4 == nil:", maps.EqualFunc(nilMap4, nil, func(v1, v2 int) bool { return true })) // Output: true
	// maps.EqualFunc(m4, m5, nil) // Panics

	// 7. Keys
	fmt.Println("\n--- Keys ---")
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	keys := maps.Keys(m)
	fmt.Println("Keys of Map:", keys) // Output: [a b c] (order indeterminate)

	// Edge Case: Keys with nil map
	var nilMap5 map[string]int
	fmt.Println("Keys of Nil Map:", maps.Keys(nilMap5)) // Output: []

	// 8. Values
	fmt.Println("\n--- Values ---")
	values := maps.Values(m)
	fmt.Println("Values of Map:", values) // Output: [1 2 3] (order indeterminate)

	// Edge Case: Values with nil map
	var nilMap6 map[string]int
	fmt.Println("Values of Nil Map:", maps.Values(nilMap6)) // Output: []
}
