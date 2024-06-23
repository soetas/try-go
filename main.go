package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

var logger = Logger{
	Scope:    "local",
	Level:    "Debug",
	Filename: "system.log",
}

func main() {
	user := struct {
		account string
		email   string
	}{"Derek Olson", "fivasem@tepnok.es"}

	Log(true, "hi,golang", 78.2)

	fmt.Printf("%v %f\n%v\n%+v\n%#v\n", math.Pi, 67.00218918, user, user, &user)
	fmt.Println(Bin(255), Chr(97))

	var data bytes = []byte("hello,world")
	var minute str = "7"

	fmt.Println(data.Decode(), minute.PadStart(2, "0"))

	logger.Info("", "start running ...")
	logger.Info("", "detection system environment ...")

	// phone := Input("enter your phone: ")
	// fmt.Println(phone)

	pwd, _ := os.Getwd()

	Tree(pwd, []string{".git"})

	Exec("")

	fmt.Println(GetGoPath(), Now())

	SetTimeout(func() {
		fmt.Println(Now())
	}, 3)

	SetInterval(func(t time.Time) {
		fmt.Println(t)
	}, 5)

}
