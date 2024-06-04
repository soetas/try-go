package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type any interface{}

type Address struct {
	Country string
	City    string
}

func (address *Address) SetCountry(country string) {
	address.Country = country
}

type User struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
	Address
}

func (user *User) Print() {
	fmt.Printf("account: %s; passwd: %s\n", user.Account, user.Passwd)
}

func CatchPanic() {
	msg := recover()

	switch msg.(type) {
	case string:
		fmt.Printf("message: %s\n", msg)
	case error:
		fmt.Printf("error: %s\n", msg)
	case int:
		fmt.Printf("error code: %d\n", msg)
	default:
		fmt.Println("unknown error")
	}
}

func Max(values ...int) int {
	if len(values) == 0 {
		panic("values is empty!")
	}

	maxValue := values[0]

	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func Example() {
	defer CatchPanic()

	hasPanic := true

	programmer := User{
		Account: "Erik Foster",
		Passwd:  "e0f7d16e-1637-5961-8a14-4b40a652c6a7",
		Address: Address{
			Country: "Romania",
		},
	}

	programmer.SetCountry("Kyrgyzstan")

	fmt.Println(len("中国"), programmer, programmer.Country)

	programmer.Print()

	p := new(int)

	var parr *[3]int

	fmt.Printf("%v %v %v %T %T\n", p, *p, p == nil, [...]*int{new(int), new(int)}, parr)

	if json, err := json.Marshal(programmer); err == nil {
		fmt.Println(string(json))
	} else {
		fmt.Println(err.Error())
	}

	stats, _ := json.Marshal(map[string]any{
		"stars":        32,
		"repositories": 30,
		"projects":     10,
	})

	var readme map[string]any

	json.Unmarshal(stats, &readme)

	fmt.Println(string(stats), readme, Max(67, 10, 78, 56, 22, 50))

	if hasPanic {
		panic(errors.New("oh, no"))
	}

	editors := make([]string, 5)

	slice := make([]string, 3)
	button := make(map[string]any)

	slice[0] = "atom"

	slice = append(slice, "vscode")

	button["size"] = "small"
	button["color"] = "primary"
	button["round"] = true

	delete(button, "color")

	fmt.Printf("%v %d %d\n", slice, len(slice), cap(slice))
	fmt.Println(button)

	copy(editors, slice)

	fmt.Println(editors)

	slice[0] = "vim"

	fmt.Println(editors)

	user := new(User)

	fmt.Printf("%T %s\n", user, reflect.TypeOf(user))
}
