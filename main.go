package main

import "github.com/soetas/webgo/rand"

const maxSize = 10

func main() {
	var (
		pointer *int
		array   [maxSize]int
		slice   []int
		dict    map[string]interface{}
	)

	Printf("%v %v %v %v", pointer, array, slice, dict)
	Printf("%v", Reversed([]any{6, true, 81.2, 1 - 8i, "Ophelia Quinn"}))

	lst := List{67, 19, 45, 99, 10, 55, 80}

	Printf("%v", lst)

	lst.Reverse()

	Printf("%v", lst)

	Printf("%v", rand.Color())

}
