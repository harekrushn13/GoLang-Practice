package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	fp := &f

	slType := reflect.TypeOf(sl)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	examiner(slType, 0)
	examiner(gType, 0)
	examiner(grpType, 0)
	examiner(fType, 0)
	examiner(fpType, 0)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type : ", t.Name(), " Kind : ", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, " Name : ", f.Name, " Type : ", f.Type.Name(), " Kind : ", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), " Tag : ", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), " tag1 : ", f.Tag.Get("tag1"), " tag2 : ", f.Tag.Get("tag2"))
			}
		}
	}
}
