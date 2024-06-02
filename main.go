package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	_ "unsafe"
)

func Between(value, min, max int) bool {
	fmt.Println("Between func called ...")

	if value >= min && value < max {
		return true
	} else {
		return false
	}
}

type User struct {
	Account string `required:"true" maxLength:"8"`
	Email   string ``
}

const (
	RECT = iota
	CIRCLE
	TRIANGLE
	ELLIPSE
	RHOMBOID
)

func TypeOf(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func recoverPanic() {
	if res := recover(); res != nil {
		if err, ok := res.(error); ok {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	defer recoverPanic()

	t := reflect.TypeOf(User{})

	if field, ok := t.FieldByName("Account"); ok {
		fmt.Printf("%T\n", field.Tag)
	} else {
		panic("unknown field")
	}

	score := 80.1

	fmt.Println(true || Between(90, 10, 90), math.Pow(math.Sqrt(score), 2) == score)

	shape := RECT

	switch shape {
	case RECT, RHOMBOID, TRIANGLE:
		fmt.Printf("")
	case CIRCLE, ELLIPSE:
		fmt.Printf("")
	default:
		fmt.Printf("")
	}

	screenWidth := 980

	switch {
	case screenWidth < 576:
		fmt.Printf("x-small")
	case screenWidth < 768:
		fmt.Printf("small")
	case screenWidth < 992:
		fmt.Printf("medium")
	case screenWidth < 1200:
		fmt.Printf("large")
	case screenWidth < 1400:
		fmt.Printf("extra large")
	default:
		fmt.Printf("extra extra large")
	}

	bytes := []byte("hi,golang")

	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%c\n", bytes[i])
	}

	for _, ch := range bytes {
		fmt.Printf("%c\n", ch)
	}

	fmt.Printf("%v %T\n", fmt.Errorf(""), fmt.Println)

	func() {
		fmt.Printf("")
	}()

	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	fmt.Println(min(61, 90))

}
