package main

import (
	"fmt"
	"strconv"
)

var Version = "1.16.6"

var (
	username string = ""
	passwd   string = ""
	email    string = ""
)

const (
	_ = iota + 1
	GREEN
	BLACK
	BLUE = iota
	TOMATO
	PINK
)

const (
	_      = iota
	SECOND = iota
	MINUTE = 1<<(6*(iota-1)) - 4*(iota-1)
	HOUR
)

func Sorted(values *[5]int) {
	for i := 0; i < len(values); i++ {
		for j := i; j < len(values); j++ {
			if values[i] > values[j] {
				tmp := values[i]
				values[i] = values[j]
				values[j] = tmp
			}
		}
	}
}

func Contains(m map[string]Repo, value string) bool {
	if _, exist := m[value]; exist {
		return true
	} else {
		return false
	}
}

type Author struct {
	Name  string
	Email string
	Repo
}

type Repo struct {
	Name  string
	Stars int
	Forks int
}

func main() {
	Version := "1.0.1"

	var baseURL = "http://localhost:8080"

	var score float32 = 78.9

	fmt.Println("hi, golang!", Version)
	fmt.Printf("%T %s %s %s %c\n", Version, username, passwd, email, 42)
	fmt.Println(baseURL, int(score), strconv.Itoa(51))

	var isOk bool
	var vector complex64 = -8

	fmt.Printf("%v %T %v %T\n", isOk, isOk, vector, vector)
	fmt.Println(0b1011, 0b1001)
	fmt.Printf("%T %v %v %v\n", 0b100, imag(vector), real((vector)), complex(0, 0))

	greet := "hi,李焕英"

	fmt.Printf("%v\n%v\n", []byte(greet), []rune(greet))
	fmt.Println(greet+"!", TOMATO, SECOND, MINUTE, HOUR)

	languages := [...]string{5: "go"}

	fmt.Println(len(languages))

	players := []string{}
	points := make([]complex64, 3, 5)

	points = append(points, []complex64{5i, 2}...)

	fmt.Println(points, players == nil)

	repos := map[string]Repo{
		"#1": {
			Name:  "yi",
			Stars: 12,
			Forks: 0,
		},
		"#2": {"100s", 67201, 10},
		"#3": {},
	}

	delete(repos, "#3")

	fmt.Println(repos["#2"], Contains(repos, "#3"))

	browser := struct {
		name    string
		engine  string
		version string
	}{"chrome", "v8", "110.0.56"}

	fmt.Println(browser)

	author := Author{
		Name:  "soetas",
		Email: "",
		Repo:  Repo{"jstdlib", 18278, 45},
	}

	fmt.Println(author.Name, author.Repo.Name)

}
