package main

import (
	"fmt"
	"log"
)

func convert(value, scale int) string {
	switch scale {
	case 2:
		return fmt.Sprintf("0b%b", value)
	case 8:
		return fmt.Sprintf("0o%o", value)
	case 16:
		return fmt.Sprintf("0x%x", value)
	default:
		panic("invalidate scale!")
	}
}

func upperCase(value string) string {
	codes := []rune(value)

	for idx, code := range codes {
		if code >= 97 && code <= 122 {
			codes[idx] -= 32
		}
	}

	return string(codes)
}

func Input() {
	fmt.Scanln()
}

func hasKey(m map[string]interface{}, key string) bool {
	if _, exist := m[key]; exist {
		return true
	} else {
		return false
	}
}

type number = float64

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		username string = "Randall Lloyd"
		email    string = "hi@vu.nf"
	)

	fmt.Printf("<%s, %s>\n", username, email)
	fmt.Println(Author, convert(12, 2))
	fmt.Printf("%T %U %s %p %v\n", '😃', '🙂', upperCase("hi, golang!"), new(int), *new(float64))
	fmt.Printf("%5.2f %s\n", 3.1415926, `\n\t\r`)

	// var rate float64

	// if _, err := fmt.Scanln(&rate); err == nil {
	// 	fmt.Printf("%v %T\n", rate, rate)
	// }

	fmt.Println(func(x, y int) (max int) {
		if x > y {
			max = x
		} else {
			max = y
		}
		return
	}(67, 99))

	players := [...]string{"Robert Sanchez", "Martin Powell", "Lina Blake"}

	points := [5][2]float64{
		{39.67, 38.04},
		{95.63, 40.30},
		{41.57, 12.54},
	}

	whiteList := []string{5: ""}

	whiteList = append(whiteList, []string{"", ""}...)

	fmt.Printf("%T %v %d %d\n", players, points, len(whiteList), cap(whiteList))

	for idx, char := range "极限挑战" {
		fmt.Printf("%d => %c\n", idx, char)
	}

	config := make(map[string]interface{})

	config["name"] = ""
	config["license"] = ""
	config["version"] = ""

	delete(config, "name")

	fmt.Println(config, hasKey(config, "name"), hasKey(config, "author"))

	var port number = 9527

	fmt.Printf("%T\n", port)

}

func init() {
	fmt.Println("main::init called ... ")
}
