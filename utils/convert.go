package utils

import (
	"fmt"
	"strconv"
)

func Str(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case float64:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func ParseInt(s string) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	} else {
		panic("")
	}
}
