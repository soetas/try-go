package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Address struct {
	Country string
	City    string
	Street  string `json:"-"`
}

type User struct {
	Account string `json:"account"`
	Email   string `json:"email,omitempty"`
	Address Address
}

var (
	ErrorValue = errors.New("error: value is wrong")
	ErrorType  = fmt.Errorf("")
)

func DivMod(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, ErrorValue
	}
	return x / y, x % y, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	html := `
	<html>
		<head></head>
		<body></body>
	</html>
	`

	if json, err := json.Marshal(html); err == nil {
		fmt.Println(string(json))
	}

	if json, err := json.Marshal(nil); err == nil {
		fmt.Println(string(json))
	}

	userData, _ := json.Marshal(&User{
		Account: "",
		Email:   "",
		Address: Address{
			Country: "",
			City:    "",
			Street:  "",
		},
	})

	buffer := new(bytes.Buffer)

	if err := json.Indent(buffer, userData, "", "  "); err == nil {
		fmt.Println(buffer.String())
	}

	_, _, err := DivMod(10, 0)

	fmt.Println(errors.Is(err, ErrorValue))

}
