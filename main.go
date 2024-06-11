package main

import (
	"fmt"
	pattern "github.com/soetas/webgo/design-pattern"
)

func main() {
	widget := pattern.CreateWidget("title")

	fmt.Printf("%#v\n", widget)

}
