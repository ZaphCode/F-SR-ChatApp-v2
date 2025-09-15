package utils

import (
	"fmt"
	"reflect"
)

const blue = "\033[34m"
const reset = "\033[0m"

func PrettyPrint(data interface{}) {
	fmt.Printf("\n%s", blue)
	val := reflect.ValueOf(data)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		fmt.Printf("%s%T [\n", blue, data)
		for i := range val.Len() {
			fmt.Printf("  %d: {\n", i)
			printStructFields(val.Index(i).Interface())
			fmt.Println("  }")
		}
		fmt.Println("]")
	case reflect.Struct:
		fmt.Printf("%s%T {\n", blue, data)
		printStructFields(data)
		fmt.Println("}")
	default:
		fmt.Printf("%s(%T) %v\n", blue, data, data)
	}
	fmt.Print(reset)
}

func printStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf("\t%s: %v\n", field.Name, value.Interface())
	}
}
