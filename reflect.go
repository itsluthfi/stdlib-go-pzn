package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Address string
}

type CustWithTag struct {
	Name    string `required:"true" max:"10"` // struct tag biasanya buat validasi, kayak bikin model di expressjs
	Address string `required:"true" max:"10"`
	ID      int    `required:"true" max:"10"`
}

// interface{} = any
func readField(value interface{}) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "has type", valueField.Type)
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	fmt.Println(t.Name())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			fmt.Println(data)
			result = data != ""
			return result
		}
	}
	return result
}

func main() {
	// readField(Person{"Luthfi", "YK"})
	isValid(CustWithTag{"Izzuddin", "YK", 482})
}
