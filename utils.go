package main

func Typeof(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	case float64:
		return "float32"
	default:
		return "unknown"
	}
}
