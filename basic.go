package main

import (
	"fmt"
	"os"
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

func (m *Matrix) ToArray() {}

func Hashable(obj any) bool {
	switch obj.(type) {
	case []int, []float64, []string:
		return false
	default:
		return true
	}
}

func HasKey(m map[string]any, k string) bool {
	if _, exist := m[k]; exist {
		return true
	} else {
		return false
	}
}

func IsNil(obj any) bool {
	return obj == nil
}

func IsBool(v any) bool {
	if _, ok := v.(bool); ok {
		return true
	} else {
		return false
	}
}

func String(v any) string {
	switch val := v.(type) {
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case float64:
		return strconv.FormatFloat(val, 'f', 6, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Hex(v int64) string {
	return "0x" + strconv.FormatInt(v, 16)
}

func ParseInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(v)
}

func ParseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)

	if err != nil {
		panic(err)
	}

	return v
}

func Create(filename string) *os.File {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte(""))
	file.WriteString("")
	file.WriteAt([]byte(""), int64(os.SEEK_CUR))

	return file
}

func Open(filename string) {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file.Close()

	block := make([]byte, 1024)

	file.Read(block)
	file.ReadAt(block, int64(os.SEEK_CUR))

}
