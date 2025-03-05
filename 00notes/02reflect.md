Here's your comprehensive markdown file for GitHub, covering Go reflection in detail, including edge cases and best practices.

---

# **Go Reflection - In-Depth Guide**

Reflection in Go provides the ability to inspect and manipulate types, values, and functions at runtime. This is useful when dealing with generic data structures, serialization, and dynamic type handling.

---

## **1. Overview of Reflection**
Reflection in Go revolves around three key concepts:

- **Types (`reflect.Type`)** – Represents the type of a variable.
- **Values (`reflect.Value`)** – Represents the actual data.
- **Kinds (`reflect.Kind`)** – Represents the underlying structure (e.g., struct, slice, map).

Reflection is implemented in the `reflect` package.

```go
import "reflect"
```

---

## **2. Understanding `reflect.Type`**
The function `reflect.TypeOf(variable)` returns a `reflect.Type` instance, allowing type inspection.

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	var text string = "Hello"

	fmt.Println(reflect.TypeOf(num))  // int
	fmt.Println(reflect.TypeOf(text)) // string
}
```

### **Methods of `reflect.Type`**
- **`Name()`** – Returns the type name. Some types (like slices, pointers) return an empty string.
- **`Kind()`** – Returns the fundamental kind (`struct`, `slice`, `map`, etc.).

```go
type Foo struct{}

func main() {
	t := reflect.TypeOf(Foo{})
	fmt.Println("Type:", t.Name())   // Foo
	fmt.Println("Kind:", t.Kind())   // struct
}
```

- **Type vs. Kind:**  
  If `Foo` is a struct, `Type` is `"Foo"`, and `Kind` is `"struct"`.

---

## **3. Inspecting Types Recursively**
A recursive function to analyze any type:

```go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type:", t.Name(), "Kind:", t.Kind())

	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "Name:", f.Name, "Type:", f.Type.Name(), "Kind:", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag:", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1:", f.Tag.Get("tag1"), "tag2:", f.Tag.Get("tag2"))
			}
		}
	}
}

type Sample struct {
	A int    `tag1:"FirstTag" tag2:"SecondTag"`
	B string
}

func main() {
	examiner(reflect.TypeOf(Sample{}), 0)
}
```

---

## **4. Understanding `reflect.Value`**
The function `reflect.ValueOf(variable)` provides access to the actual value.

```go
var num int = 100
v := reflect.ValueOf(num)
fmt.Println(v.Int()) // 100
```

### **Modifying Values**
To modify values, use a pointer:

```go
func modifyVariable(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v.Elem().SetInt(200)
	}
}

func main() {
	num := 100
	modifyVariable(&num)
	fmt.Println(num) // 200
}
```

- **Modification requires a pointer**; otherwise, it panics.

---

## **5. Struct Field Reflection & Tags**
Reflection allows reading struct field details, including tags.

```go
type Foo struct {
	A int    `json:"field_a"`
	B string `json:"field_b"`
}

func main() {
	x := Foo{A: 10, B: "Hello"}
	t := reflect.TypeOf(x)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println("Field:", f.Name, "Type:", f.Type, "Tag:", f.Tag.Get("json"))
	}
}
```

**Output:**
```
Field: A Type: int Tag: field_a
Field: B Type: string Tag: field_b
```

---

## **6. Creating and Modifying Structs at Runtime**
Reflection allows creating new struct instances dynamically.

```go
fields := []reflect.StructField{
	{Name: "X", Type: reflect.TypeOf(int(0))},
	{Name: "Y", Type: reflect.TypeOf(string(""))},
}

structType := reflect.StructOf(fields)
newStruct := reflect.New(structType).Elem()

newStruct.Field(0).SetInt(100)
newStruct.Field(1).SetString("Hello")

fmt.Println(newStruct.Interface()) // {100 Hello}
```

---

## **7. Creating Functions at Runtime**
Dynamic function creation using `reflect.MakeFunc`:

```go
func dynamicFunc(in []reflect.Value) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(in[0].Int() * 2)}
}

func main() {
	funcType := reflect.FuncOf([]reflect.Type{reflect.TypeOf(0)}, []reflect.Type{reflect.TypeOf(0)}, false)
	newFunc := reflect.MakeFunc(funcType, dynamicFunc)

	result := newFunc.Call([]reflect.Value{reflect.ValueOf(10)})
	fmt.Println(result[0].Int()) // 20
}
```

---

## **8. Edge Cases & Pitfalls**
### **8.1. Reflection Performance Overhead**
Reflection is **slower** than direct method calls. Use it **only when necessary**.

### **8.2. Accessing Unexported Struct Fields**
Unexported fields cannot be accessed directly:

```go
type Foo struct {
	privateField int
}

func main() {
	f := Foo{privateField: 42}
	v := reflect.ValueOf(f)
	fmt.Println(v.FieldByName("privateField").CanSet()) // false
}
```

**Solution:** Use `reflect.ValueOf(&f).Elem()` to modify fields.

---

## **9. Reflection Best Practices**
1. **Use Reflection Sparingly** – It can lead to performance issues and complex code.
2. **Check `Kind()` Before Using Reflection Methods** – Avoid panics.
3. **Avoid Modifying Values Unless Necessary** – Use normal type assertions when possible.
4. **Use Struct Tags Properly** – Metadata in tags should follow proper conventions.
5. **Benchmark Performance** – Reflection should be used only when required.

---