package main

import (
	"fmt"
	"reflect"
)

const PI = 3.1415926

func Swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	greet := "hi, go!"

	var age uint8 = 25
	var address string
	var score float64
	var comparator func(int, int) bool

	age++

	fmt.Println(greet, reflect.TypeOf(age), address == "", score == 0)
	fmt.Println(comparator == nil)
	fmt.Printf("%T\n", .0)

	pointerOfAge := &age

	fmt.Printf("%p %v\n", &age, pointerOfAge)

	// block scope
	{
		var x, y int = 78, 45

		fmt.Printf("before: x=%d, y=%d\n", x, y)

		Swap(&x, &y)

		fmt.Printf("after: x=%d, y=%d\n", x, y)
	}

	var origin complex64
	var agree bool

	fmt.Println(origin, agree, "\"Christine"+" Cain\"", uint8(12.0))
	fmt.Printf("%.2f\n", PI)

	if true {
		fmt.Printf("")
	} else {
		fmt.Printf("")
	}

	switch {
	case age < 18:
		fmt.Printf("study hard and make progress every day\n")
	case age < 30:
		fmt.Printf("come on, bitch\n")
		fallthrough
	case age < 50:
		fmt.Printf("learn to be grateful\n")
	default:
		fmt.Printf("borrow another 500 years from heaven\n")
	}

	for _, ch := range "歌手2024" {
		fmt.Printf("%c\n", ch)
	}

	// fmt.Println(ch)
	fmt.Printf("%T\n", fmt.Printf)

}
