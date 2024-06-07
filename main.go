package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/golang-lru"
	"github.com/soetas/webgo/json"
	"github.com/soetas/webgo/random"
	_ "github.com/soetas/webgo/this"
	"log"
	"os"
	"time"
)

type any = interface{}

var (
	ErrorZeroDivision = errors.New("divisor cannot be 0")
	ErrorType         = fmt.Errorf("type mismatch")
)

type ErrorIndex struct {
	Message string
}

func (err *ErrorIndex) String() string {
	return err.Message
}

func SingleReturnValue(f func(...any) (any, error)) {}

func Divmod(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, ErrorZeroDivision
	} else {
		return x / y, x % y, nil
	}
}

func Quit() {
	os.Exit(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	user := map[string]any{
		"account": "Elmer Owens",
		"email":   "dilwel@becfozzog.to",
	}

	fmt.Println(random.Choice([]any{"pink", "skyblue", "tomato"}))
	fmt.Println(json.Stringify(user))

	if div, mod, err := Divmod(10, 3); err == nil {
		fmt.Println(div, mod)
	} else {
		if errors.Is(err, ErrorZeroDivision) {
			panic(err)
		}
	}

	fmt.Printf("%p %p\n", errors.New(""), errors.New(""))
	fmt.Println(time.Now().Add(time.Minute * 10).Format("2006/01/02 15:04:05"))

	cache, _ := lru.New(120)

	cache.Add("jim", 67.1)

}
