package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var (
	ErrKey   = errors.New("")
	ErrIndex = fmt.Errorf("")
)

type ZeroDivisionError struct{}

func (t *ZeroDivisionError) Error() string {
	return "division by zero"
}

func Div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, &ZeroDivisionError{}
	} else {
		return x / y, nil
	}
}

func Main11() {
	defer CatchPanic()

	if files, err := ioutil.ReadDir("templates"); err == nil {
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		os.Exit(1)
	}

	fmt.Printf("%T\n", errors.New(""))

	_, err := Div(10, 0)

	if _, ok := err.(*ZeroDivisionError); ok {
		fmt.Printf("%+v %T\n", err, err)
	}

	scores := []int{86, 61, 99, 59, 51, 68, 89}

	sort.Ints(scores)

	fmt.Println(scores, sort.SearchInts(scores, 77))

	users := []map[string]any{
		{"account": "Amanda Tucker", "email": "ebvil@obiovuzu.tk", "points": 9},
		{"account": "Celia Griffith", "email": "edga@jeonaki.es", "points": 40},
		{"account": "Katie Stephens", "email": "olennok@akicoumi.zw", "points": 16},
		{"account": "Craig Larson", "email": "uhejah@zahav.bt", "points": 15},
		{"account": "Willie Wood", "email": "he@behme.np", "points": 46},
	}

	sort.Slice(users, func(i, j int) bool {
		x, _ := users[i]["points"].(int)
		y, _ := users[j]["points"].(int)
		return x <= y
	})

	fmt.Println(users)

	fmt.Printf("%d %v\n", len(os.Args), os.Args)

	open := *(flag.Bool("open", false, ""))
	host := *(flag.String("host", "", ""))
	port := *(flag.Int("port", 5713, ""))

	flag.Parse()

	fmt.Printf("[server]\nopen=%v\nhost=%s\nport=%d\n", open, host, port)
}
