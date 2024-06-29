package main

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

type (
	Number int
	Float  float64
	Char   rune
)

var keywords = []string{"package", "import", "var", "func", "type"}

const defaultEncoding = "utf-8"

const (
	Ayu = iota
	Material
	Dainty
	Dracula
	Horizon
)

func Printf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func Typeof(obj any) string {
	switch obj.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case string:
		return "string"
	case byte:
		return "byte"
	case rune:
		return "rune"
	case complex128:
		return "complex128"
	default:
		return fmt.Sprintf("%T", obj)
	}
}

func Swap(x, y *string) {
	*x, *y = *y, *x
}

func Sizeof(obj any) {}

func Scanf(prompt string) string {
	var enter string

	fmt.Print(prompt)
	fmt.Scan(&enter)

	return enter
}

func Int(value string) int {
	res, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(res)
}

func Float64(value string) float64 {
	res, err := strconv.ParseFloat(value, 64)

	if err != nil {
		panic(err)
	}

	return res
}

func Recover() {
	if res := recover(); res == nil {
		if err, ok := res.(error); ok {
			fmt.Println(err.Error())
		}
	}
}
