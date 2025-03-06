package main

import (
	"fmt"
	"time"
)

func main() {

	defer func() {

		if r := recover(); r != nil {

			fmt.Println("panic occurrred")

			return
		}
	}()

	value1 := make(map[string]string)

	value1["key"] = "value"

	for index := 0; index < 100; index++ {

		go func(values map[string]string) {

			defer func() {

				if r := recover(); r != nil {

					fmt.Println("panic occurrred")

					return
				}
			}()

			//value1 = value1[:500]

			for index := 0; index < 100; index++ {

				values["key"] = "value"

			}

			return

		}(value1)

		go func(values map[string]string) {

			defer func() {

				if r := recover(); r != nil {

					fmt.Println("panic occurrred")

					return
				}
			}()

			//value1 = value1[:500]

			for index := 0; index < 100; index++ {

				if _, ok := values["key"]; ok {

					fmt.Println("key is present")
				}

			}

			return

		}(value1)

		go func(values map[string]string) {

			defer func() {

				if r := recover(); r != nil {

					fmt.Println("panic occurrred")

					return
				}
			}()

			for index := 0; index < 100; index++ {

				if _, ok := values["key"]; ok {

					fmt.Println("key is present")
				}

			}
			return

		}(value1)

	}

	time.Sleep(time.Second * 100)

	fmt.Println("i am here")

}

func Abc(value1, value3 []string, value2 []int8) bool {

	return true

}
