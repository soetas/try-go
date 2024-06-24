package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

var enableLog = false

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	if enableLog {
		logFile, err := os.OpenFile("system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			log.Fatal(err.Error())
		}

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetPrefix("[Output] ")
		log.SetOutput(logFile)

		log.Print("start install ...")
		log.Printf("fetch latest version of packages ...")
		log.Println("waiting for install ...")
	}

	if divVal, _, err := Divmod(19, 3); err == nil {
		fmt.Println(divVal)
	}

	fmt.Printf("%s %t\n", bytes.ToUpper([]byte("hi,李焕英")),
		bytes.EqualFold([]byte("hellO,World"), []byte("HEllo,worlD")))

	url := "https://fanyi.youdao.com/#/"
	hostname := bytes.Trim(bytes.TrimPrefix([]byte(url), []byte("https:")), "/#")

	fmt.Printf("%s\n", hostname)
	fmt.Printf("%s\n", bytes.Repeat([]byte{'='}, 80))
	fmt.Println(bytes.Runes([]byte("hi,李焕英")))

	users := make([]string, 0, 5)

	users = append(users, []string{"Chris May", "Polly Hubbard"}...)

	fmt.Println(len(users), cap(users), users)

	res := Response{
		Code:    2000,
		Message: "success",
		Data: []map[string]interface{}{
			{
				"address": nil,
			},
			{},
			{},
		},
	}

	if result, err := json.Marshal(res); err == nil {
		fmt.Println(string(result))
	}

	scores := []float64{78, 88, 86, 90, 67}

	sort.Float64s(scores)

	fmt.Println(scores, Choice([]interface{}{"green", "navy", "tomato", "orange", "pink"}))

	// GetCommandArgs()

	program := Command{}

	program.Name("version")
	program.Name("open")
	program.Name("browser")

	program.Parse()

}
