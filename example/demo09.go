package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"
)

type StringSlice []string

func CloseTo(value, other float64) bool {
	return math.Abs(value-other) < 0.00001
}

func Hex(value int) string {
	return fmt.Sprintf("0x%08x", value)
}

func CodePointAt(str string, pos int) int {
	return 0
}

func Sum(values ...float64) (s float64) {
	for _, value := range values {
		s += value
	}
	return
}

func (n Number) ToString() string {
	return fmt.Sprintf("%v\n", n)
}

func (s StringSlice) Reverse() {}

type object = map[string]interface{}

func HasKey(obj object, key string) bool {
	if _, ok := obj[key]; ok {
		return true
	} else {
		return false
	}
}

func Counter(slice []any) map[any]int {
	result := make(map[any]int)

	for _, item := range slice {
		result[item]++
	}

	return result
}

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPost(id int, title, body string) *Post {
	return &Post{
		Id:    id,
		Title: title,
		Body:  body,
	}
}

func Main09() {
	fmt.Printf("%T %f %.2f %t\n", .0, math.E, math.Pi, CloseTo(0.1+0.2, 0.3))
	fmt.Println(Hex(141))
	fmt.Printf("%T\n", 56e1)

	count := new(big.Int)

	count.SetString("89172781672671898717826716762771782781", 10)
	count.Add(count, big.NewInt(10))

	sep := '&'
	score := 78.56

	fmt.Printf("%v %v %c %c %d\n", count, `\n\t\r `, 960, sep, 'a'-'A')
	fmt.Println(utf8.RuneCountInString("🍇"), int(score))
	fmt.Println(strconv.Itoa(65), fmt.Sprintf("%d", 65))

	if value, err := strconv.Atoi("12px"); err == nil {
		fmt.Printf("after convert: %d\n", value)
	} else {
		fmt.Println(err)
	}

	fmt.Println(Sum(3, 8, 5))

	var price Number = 4

	fmt.Println(price.ToString())

	players := [...]string{
		"Dominic Love",
		"Travis Wilson",
		"Andre Thomas",
	}

	fmt.Printf("%T %v %v\n", players, strings.TrimSpace(" Mason Watts "), strings.Join(players[:], " "))

	var countries StringSlice = StringSlice{
		"Christmas Island",
		"France",
		"St. Pierre & Miquelon",
	}

	countries.Reverse()

	fmt.Println(countries)

	user := object{
		"account": "Lucinda Castillo",
		"email":   "kejde@ellu.za",
		"address": "",
	}

	delete(user, "address")

	fmt.Println(user, HasKey(user, "address"))
	fmt.Println(Counter([]any{}))

	point := struct {
		x float64
		y float64
	}{0, 0}

	posts := []Post{
		{Id: 68, Title: "", Body: ""},
		{Id: 64, Title: "", Body: ""},
		{Id: 2, Title: "", Body: ""},
	}

	fmt.Printf("%+v\n", point)
	fmt.Println(posts, NewPost(0, "", ""))

}
