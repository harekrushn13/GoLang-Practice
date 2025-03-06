```bash
package main

import "fmt"

func main() {
	var a = []string{
		"C++",
		"Go",
		"Java",
		"TypeScript",
		"TypeScript2",
	}

	fmt.Println("Array:", a, len(a), cap(a))
	a = a[0:2] 
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[:3]  
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[2:]  
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[0:]
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[:3] 
	fmt.Println("Array:", a, len(a), cap(a))
}
```

### Output
```bash
Array: [C++ Go Java TypeScript TypeScript2] 5 5
Array: [C++ Go] 2 5
Array: [C++ Go Java] 3 5
Array: [Java] 1 3
Array: [Java] 1 3
Array: [Java TypeScript TypeScript2] 3 3
```

```bash
package main

import "fmt"

func main() {
	
	s1 := make([]string, 4)
	fmt.Println("Array:", s1, len(s1), cap(s1))
	_ = copy(s1, []string{"a", "b", "c", "d"})
	fmt.Println("Array:", s1, len(s1), cap(s1))
	s1 = append(s1, "e", "f")
	fmt.Println("Array:", s1, len(s1), cap(s1))
	s1 = append(s1, "g", "h")
	fmt.Println("Array:", s1, len(s1), cap(s1))
	s1 = append(s1, "i")
	fmt.Println("Array:", s1, len(s1), cap(s1))
}
```

### Output
```bash
Array: [   ] 4 4
Array: [a b c d] 4 4
Array: [a b c d e f] 6 8
Array: [a b c d e f g h] 8 8
Array: [a b c d e f g h i] 9 16
```