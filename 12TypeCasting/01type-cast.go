package main

import (
	"fmt"
)

func main() {
	// 1. Integer to Integer
	var a int = 123456
	var b int8 = int8(a) // Edge case: may cause overflow if a is too large
	fmt.Printf("int to int8: %d -> %d\n", a, b)

	// Edge case: large integer to small type, overflow.
	var largeInt int64 = 999999999999
	var smallInt int8 = int8(largeInt) // Edge case: overflow happens here
	fmt.Printf("int64 to int8 (overflow): %d -> %d\n", largeInt, smallInt)

	// 2. Integer to Floating-point
	var x int = 10
	var y float64 = float64(x) // valid conversion
	fmt.Printf("int to float64: %d -> %f\n", x, y)

	// Edge case: loss of precision when converting float64 to int.
	var z float64 = 123.987
	var n int = int(z) // Edge case: truncates the decimal part
	fmt.Printf("float64 to int (loss of precision): %f -> %d\n", z, n)

	// 3. Floating-point to Integer
	var f1 float32 = 3.14
	var i1 int = int(f1) // valid conversion, truncates decimals
	fmt.Printf("float32 to int: %f -> %d\n", f1, i1)

	// Edge case: large float to smaller integer
	var f2 float64 = 1234567890.123456
	var i2 int32 = int32(f2) // Edge case: loss of precision and truncation
	fmt.Printf("float64 to int32 (loss of precision): %f -> %d\n", f2, i2)

	// 4. Integer to Boolean
	// Go does NOT allow direct conversion from integer to boolean.
	// Uncommenting the next line will cause a compile-time error:
	//var b1 bool = bool(a) // ERROR: cannot convert int to bool

	// 5. Boolean to Integer
	var b2 bool = true
	var i3 int = 0
	if b2 {
		i3 = 1
	}
	fmt.Printf("bool to int: %v -> %d\n", b2, i3)

	// 6. String to Rune (Character)
	var str string = "hello"
	var r rune = []rune(str)[0] // Accessing first character of string
	fmt.Printf("string to rune: %s -> %c\n", str, r)

	// 7. Rune to String
	var r2 rune = 'A'
	var str2 string = string(r2) // Convert rune to string
	fmt.Printf("rune to string: %c -> %s\n", r2, str2)

	// 8. Array to Slice
	var arr = [3]int{1, 2, 3}
	var sl = arr[:] // Array to Slice conversion
	fmt.Printf("array to slice: %v -> %v\n", arr, sl)

	// 9. Slice to Array (not directly possible)
	// Edge case: Conversion of slice to array is not allowed directly
	// Uncommenting the next line will cause a compile-time error:
	//var arr2 [3]int = sl // ERROR: cannot convert slice to array

	// 10. Pointer to Pointer (same type)
	var p *int = &x
	var q *int = p // valid conversion
	fmt.Printf("pointer to pointer: %v -> %v\n", p, q)

	// Edge case: Pointer to incompatible type (compile-time error)
	// Uncommenting the next line will cause a compile-time error:
	//var r *float64 = p // ERROR: cannot assign *int to *float64

	// 11. Type Assertion (Interface to Concrete Type)
	var i4 interface{} = 42
	var i5 int = i4.(int) // valid type assertion
	fmt.Printf("interface to concrete type: %v -> %d\n", i4, i5)

	// Edge case: Type assertion fails (panic if type doesn't match)
	// Uncommenting the next line will cause a runtime panic:
	//var i6 float64 = i4.(float64) // ERROR: panic: interface conversion: interface {} is int, not float64

	// Safe Type Assertion with "comma ok"
	if v, ok := i4.(float64); ok {
		fmt.Println("type assertion succeeded:", v)
	} else {
		fmt.Println("type assertion failed: i4 is not a float64")
	}

	// 12. Struct Type Conversion
	type StructA struct {
		X int
	}
	type StructB struct {
		Y int
	}

	//var a1 StructA = StructA{X: 10}
	// Edge case: Structs of different types cannot be directly converted.
	// Uncommenting the next line will cause a compile-time error:
	//var b1 StructB = a1 // ERROR: cannot use a1 (type StructA) as type StructB in assignment

	// Custom type conversion
	type MyInt int
	var m1 MyInt = 5
	var n1 int = int(m1) // valid conversion from MyInt to int
	fmt.Printf("custom type MyInt to int: %d -> %d\n", m1, n1)

	// 13. Custom type to another custom type with the same underlying type
	type MyInt2 int
	var m2 MyInt = 10
	var n2 MyInt2 = MyInt2(m2) // valid conversion
	fmt.Printf("custom type MyInt to MyInt2: %d -> %d\n", m2, n2)

	// Edge case: invalid conversion between different custom types with the same underlying type
	// Uncommenting the next line will cause a compile-time error:
	//var n3 MyInt2 = MyInt(m2) // ERROR: cannot convert MyInt to MyInt2
}
