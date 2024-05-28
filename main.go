package main

import (
	"fmt"
	"stdlib/random"
	"unicode"
)

var (
	version = "1.16.6"
	module  = true
	arch    = "amd64"
)

const PI float64 = 3.14

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)

var swap = func(x, y int) (int, int) {
	return y, x
}

func main() {
	var account string
	var points int
	var rate float32
	var good bool
	var gender rune = '👨'
	var pointer uintptr

	fmt.Printf("%s %d %.2f %t %c\n", account, points, rate, good, gender)
	fmt.Println(unicode.IsSpace('\r'))
	fmt.Printf("%x %c %c\n", gender, '\u4e25', '\U0001f468')
	fmt.Println(`\n\t\r\u\U`, len("你好，golang"))

	var template = `
	<div>

	</div>
	`

	fmt.Println(template, []rune("golang"), pointer)

	email := "hezhas@kej.pn"

	fmt.Println(email, version, module, arch, KB, GB)
	fmt.Println(10 % 3)

	var p *int = new(int)

	fmt.Printf("%v %d\n", p, *p)
	fmt.Println(swap(90, 78))

	random.Choice(67, true, "")

	if age := 18; age <= 30 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	users := []string{"Benjamin Carson", "Lenora Weaver", "Ricky Jefferson"}

	for index, user := range users {
		fmt.Printf("%d => %s\n", index, user)
	}

}
