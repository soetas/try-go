package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
}

func typeof(value interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(value).Kind())
}

func view(value interface{}) {
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		val := reflect.ValueOf(value)
		for i := 0; i < val.NumField(); i++ {
			fmt.Printf("%v\n", val.Field(i))
		}
	}
}

func main() {
	fmt.Println(typeof(1i), typeof(User{}))

	view(User{
		Name:  "Brett Curtis",
		Email: "wi@hov.in",
	})

}
