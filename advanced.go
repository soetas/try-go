package main

import (
	"fmt"
	"runtime"
	"time"
)

func parse() {
	defer func() {
		fmt.Printf("parse goroutine end ...\n")
	}()

	fmt.Printf("parse goroutine start ...\n")

	for i := 0; i < 5; i++ {
		if i >= 3 {
			runtime.Goexit()
		}
		fmt.Printf("%d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func pause() {
	for {
		break
	}
}

func main() {
	fmt.Printf("main goroutine start ...\n")

	go parse()

	pause()
}
