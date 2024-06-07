package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/soetas/webgo/utils"
)

func main() {
	url := "https://www.bilibili.com/video/BV1s341147US?p=31&spm_id_from=pageDriver"

	fmt.Println(utils.Choice([]string{"Lina Ray", "Isaac Hart", "Shawn Carson"}))
	fmt.Println(utils.Str(false), utils.ParseInt("78"))

	schema := strings.Split(url, "//")[0]

	fmt.Println(strings.Contains(url, "https"), schema)
	fmt.Println(utils.Len("你好, golang"))

	utils.Delay(1)

	fmt.Println(utils.Now("2006年01月01号 15时04分05秒"))

	utils.Warn("xx")

	println(complex(-2, 8), runtime.GOARCH, runtime.GOOS, runtime.NumCPU())

}
