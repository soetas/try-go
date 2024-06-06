package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Str(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%d", value)
	case bool:
		return fmt.Sprintf("%t", value)
	case float64:
		return fmt.Sprintf("%f", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func ParseInt(s string) (int, error) {
	if res, err := strconv.Atoi(s); err == nil {
		return res, nil
	} else {
		return 0, fmt.Errorf("")
	}
}

func Main06() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64(), Str(78.12), Str(false))

	if pixel, err := ParseInt("12.12"); err == nil {
		fmt.Printf("%v %T\n", pixel, pixel)
	}

	searchParams := "q=flutter&form=QBLH&sp=-1&lq=0&pq=flu&sc=10-3&qs=n"

	fmt.Println(strings.Contains("", ""), strings.ReplaceAll("", "", ""))
	fmt.Println(strings.Fields(""), strings.Split(searchParams, "&"), strings.Trim("", ""))

	greet := "你好，世界！"

	fmt.Println(utf8.RuneCountInString(greet), utf8.ValidString(""))

	if duration, err := time.ParseDuration(""); err == nil {
		fmt.Println(duration.String())
	}

	fileInfo, err := os.Stat("")

	fmt.Println(fileInfo.IsDir(), os.IsExist(err))

}
