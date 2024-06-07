package utils

import (
	"os"
	"time"
	"unicode/utf8"
)

func Len(s string) int {
	return utf8.RuneCountInString(s)
}

func Delay(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}

func Now(format string) string {
	return time.Now().Format(format)
}

func Quit() {
	os.Exit(0)
}
