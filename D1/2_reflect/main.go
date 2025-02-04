package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `required:"true" max:"10" min:"2"`
	Email string `required:"true" max:"10" min:"2"`
}

func main() {
	user := User{
		Name:  "Fahmi",
		Email: "fahmafsaa",
	}

	valueOfUser := reflect.ValueOf(user)

	typeOfUser := reflect.TypeOf(user)
	fmt.Println(typeOfUser)
	fmt.Println(typeOfUser.Field(0).Tag.Get("required"))
	fmt.Println(typeOfUser.Field(0).Tag.Get("max"))
	fmt.Println(typeOfUser.Field(0).Tag.Get("min"))

	if typeOfUser.Field(0).Tag.Get("required") == "true" {
		if valueOfUser.IsNil() {
			fmt.Errorf("Required")
		}
	}

	if typeOfUser.Field(0).Tag.Get("max") == "10" {
		if valueOfUser.Len() > 10 {
			fmt.Errorf("not max")
		}
	}
	var number float64 = 20.12

	value := reflect.ValueOf(number)
	fmt.Println(value.Type())
	fmt.Println(value.Kind())

	typeNumber := reflect.TypeOf(number)
	fmt.Println(typeNumber)

	fmt.Println(reflect.Float64 == value.Kind())
}
