package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	a int32
	b int64
}

func main() {
	// Basic Type Conversion

	//var i int = 42
	//ptr := unsafe.Pointer(&i)
	//iptr := (*int)(ptr)
	//fptr := (*float64)(ptr)
	//
	//fmt.Println("Original int:", i)
	//fmt.Println("Original int ptr:", ptr)
	//fmt.Println("Interpreted as int:", *iptr)
	//fmt.Println("Interpreted as float64:", *fptr)

	// Memory Address Arithmetic

	//s := MyStruct{a: 10, b: 20}
	//
	//ptr := unsafe.Pointer(&s)            // Convert struct pointer to unsafe.Pointer
	//bPtr := (*int64)(unsafe.Add(ptr, 8)) // Move 8 bytes (size of int32 + padding) to get `b`
	//
	//fmt.Println("Field b:", *bPtr) // 20

	// Converting uintptr to unsafe.Pointer

	//x := 100
	//fmt.Println(&x)
	//ptr := unsafe.Pointer(&x) // Convert *int to unsafe.Pointer
	//fmt.Printf("%p\n", ptr)
	//uintAddr := uintptr(ptr) // Convert to uintptr
	//fmt.Printf("%v\n", uintAddr)
	//newPtr := unsafe.Pointer(uintAddr) // Convert back to unsafe.Pointer
	//
	//fmt.Println(newPtr)
	//fmt.Println("Original Address of x:", (*int)(newPtr))
	//fmt.Println("Original:", *(*int)(newPtr)) // 100

	// Reading and Writing Arbitrary Memory

	//var x float64 = 10
	//ptr := unsafe.Pointer(&x) // Convert *float64 to unsafe.Pointer
	//
	//*(*float64)(ptr) = 42 // Modify memory directly
	//
	//fmt.Println("Modified x:", x) // 42

	// ----------------
	//s := "hello"
	//b := []byte(s) // Creates a new byte slice with copied data
	//b[0] = 'H'     // Modifying b does NOT affect s
	//
	//fmt.Println(s)         // "hello"
	//fmt.Println(string(b)) // "Hello"

	// ----------------
	//b := []byte{'h', 'e', 'l', 'l', 'o'}
	//s := string(b) // Creates a new string with copied data
	//
	//b[0] = 'H' // Modifying b does NOT affect s
	//
	//fmt.Println(s)         // "hello"
	//fmt.Println(string(b)) // "Hello"

	//s := "Hello"
	//ptr := (*[]byte)(unsafe.Pointer(&s))
	//fmt.Println(ptr, &(*ptr)[0])
	//*ptr = []byte(s)
	//fmt.Println(ptr, &(*ptr)[0])
	//(*ptr)[0] = 'A'
	//fmt.Println(s)

	// Zero-Copy Conversion Using unsafe
	s := "hello"
	//b := StringToBytes(s)
	b := *(*[]byte)(unsafe.Pointer(&s))

	b[0] = 'H' // Modifying b WILL affect s!

	fmt.Println(s)         // "Hello" (unsafe modification)
	fmt.Println(string(b)) // "Hello"

}

func StringToBytes(s string) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s)), len(s))
}
