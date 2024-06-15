package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

type any = interface{}

type object map[string]interface{}

func (o object) toJson() string {
	res, _ := json.Marshal(o)
	return string(res)
}

func typeof(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

func has(m map[string]interface{}, k string) bool {
	if _, ok := m[k]; ok {
		return true
	} else {
		return false
	}
}

func log(values ...interface{}) {
	fmt.Println(values...)
}

const PI = math.Pi

const (
	UNSENT = iota
	OPENED
	LOADING
	DONE
)

type HTTPServer struct {
	Hostname string
	Port     int
	Open     bool
}

var (
	ErrType = errors.New("")
)

func main() {
	defer func() {
		if res := recover(); res != nil {
			if err, ok := res.(error); ok {
				fmt.Println(err.Error())
			}
		}
	}()

	guid := "ba730025-9b81-57c4-b5d1-d765e7bb4c75"

	var (
		username string = "Dylan Gonzales"
		email    string = "fe@vek.gi"
	)

	fmt.Println(typeof(0b101), typeof((guid)), fmt.Sprintf("<%s, %s>", username, email))
	fmt.Println(unsafe.Sizeof(0), math.MaxInt8)
	fmt.Println(typeof(.5))

	cities := [...]string{3: "TJ", 5: "SH"}
	coords := [...][2]float64{
		{69.11, 34.28},
		{19.49, 41.18},
		4: {-170.43, -14.16},
	}

	fmt.Println(cities, typeof(cities), typeof(coords))

	players := []string{"Cornelia Greer", 2: "Tony Greer", "Wayne Long"}

	players = append(players, []string{"Kenneth Goodwin"}...)

	fmt.Println(players)
	fmt.Println(typeof(players), len(players), cap(players))

	headers := make(map[string]interface{})

	headers["Content-Type"] = "application/json"
	headers["Accept-Language"] = "zh-CH, en"
	headers["Host"] = "localhost"

	delete(headers, "Host")

	fmt.Printf("%#v %t\n", headers, has(headers, "Host"))

	pointerOfCities := &cities

	pointerOfCities[3] = "BJ"
	(*pointerOfCities)[5] = "SZ"

	fmt.Println(cities, (pointerOfCities))

	server := HTTPServer{}

	server.Port = 8080

	fmt.Printf("%#v %t\n", server, reflect.DeepEqual(server, HTTPServer{}))

	post := object{
		"id":    0,
		"title": "",
		"body":  "",
	}

	log(0, 0xff, post.toJson())

	var label any = 0

	label = ""

	label = []string{}

	fmt.Println(label)
}
