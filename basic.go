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

var Keywords = []string{"package", "import", "var", "func", "type"}

const DefaultEncoding = "utf-8"

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

type Slice struct {
	Start int
	End   int
	Step  int
}

func Range(start, end, step int) *Slice {
	return &Slice{
		Start: start,
		End:   end,
		Step:  step,
	}
}

type Set struct {
	Items []any
}

func NewSet(values ...any) *Set {
	return &Set{
		Items: values,
	}
}

func CreateCounter(initial int) (int, map[string]func()) {
	count := initial

	actions := map[string]func(){
		"inc":   func() { count++ },
		"dec":   func() { count-- },
		"reset": func() { count = initial },
	}

	return count, actions
}

func Reversed(s []any) []any {
	result := make([]any, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

type List []any

func (l *List) Reverse() {
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}

func (l *List) Sort() {}

type Matrix struct {
	Row  int
	Col  int
	Data []float64
}

func (m *Matrix) ToArray() {

}
