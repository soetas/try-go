package main

import (
	"fmt"
	"strings"
	"time"
)

func PadStart(str string, targetLength int, padString string) string {
	format := fmt.Sprintf("%%%s%ds", padString, targetLength)
	return fmt.Sprintf(format, str)
}

func Main08() {
	fmt.Printf("hi, golang!\n")
	fmt.Println(67/float64(3), 67%3)

	now := time.Now()
	hour, minute, second := now.Hour(), now.Minute(), now.Second()

	fmt.Printf("%02d:%02d:%02d\n", hour, minute, second)
	fmt.Println(PadStart("3", 2, "0"))

	url := "https://developer.mozilla.org/zh-CN/docs"

	if strings.HasPrefix(url, "https") {
		fmt.Printf("schema: https\n")
	}

}
