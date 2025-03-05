package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	dataBytes := ReadJson("./11reflect/data.json")

	var data any
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	dataValue := reflect.ValueOf(data)
	//generateMap(dataValue)

	mpp := parseMap(dataValue).(map[string]interface{})
	fmt.Println(mpp)

	//generateMap(reflect.ValueOf(mpp))
	fmt.Println(mpp["friends"])
}

func parseMap(dataValue reflect.Value) any {
	if dataValue.Kind() == reflect.Interface {
		dataValue = dataValue.Elem()
	}

	switch dataValue.Kind() {
	case reflect.Map:
		result := make(map[string]any)
		iter := dataValue.MapRange()
		for iter.Next() {
			key := iter.Key().Interface().(string)
			value := iter.Value()
			result[key] = parseMap(value)
		}
		return result

	case reflect.Slice:
		var result []any
		for i := 0; i < dataValue.Len(); i++ {
			result = append(result, parseMap(dataValue.Index(i)))
		}
		return result

	default:
		return dataValue.Interface()
	}
}

func generateMap(dataValue reflect.Value) {
	if dataValue.Kind() == reflect.Interface {
		dataValue = dataValue.Elem()
	}

	switch dataValue.Kind() {
	case reflect.Map:
		iter := dataValue.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()

			fmt.Printf("Key: %v\t-%v\t\t", key.Interface(), key.Kind())

			generateMap(value)
		}
	case reflect.Slice:
		for i := 0; i < dataValue.Len(); i++ {
			fmt.Printf("\nIndex %d:\n", i)
			generateMap(dataValue.Index(i))
		}
	default:
		fmt.Printf("Value: %v\t-%v\n", dataValue.Interface(), dataValue.Kind())
	}
}

func ReadJson(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
