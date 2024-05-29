package main

import (
	"fmt"
)

func Divmod(x, y int) (div, mod int) {
	div = x / y
	mod = x % y
	return
}

func Open(path string) (int, error) {
	return 0, nil
}

type Devtool struct {
	Debug bool
}

type Browser struct {
	Name    string
	Version string
	Engine  string
	Devtool
}

func (b *Browser) Update(version string) {
	b.Version = version
}

type Component interface {
	Render() string
}

type Button struct{}
type Text struct{}
type Input struct{}

func (button *Button) Render() string {
	return "button"
}

func (text *Text) Render() string {
	return "text"
}

func (input *Input) Render() string {
	return "input"
}

type Any interface{}

func Typeof(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}

func Isbool(value interface{}) bool {
	if _, ok := value.(bool); ok {
		return true
	} else {
		return false
	}
}

func main() {
	colors := [5]string{}
	langs := [...]string{2: "go", 5: "python"}

	colors[0] = "green"
	colors[3] = "tomato"

	for index, color := range colors {
		fmt.Printf("%d: %s\n", index, color)
	}

	fmt.Println(len(langs))

	slice := colors[1:4]

	slice[0] = "pink"

	fmt.Printf("%v %v %d %d\n", slice, colors, len(slice), cap(slice))

	skills := make([]string, 2)

	skills[0] = "go"
	skills[1] = "linux"

	fmt.Printf("%p \n", skills)

	skills = append(skills, "docker")

	fmt.Printf("%p %d %T\n", skills, cap(skills), skills)

	if cap(skills) != 2 {
		fmt.Println("")
	} else {
		fmt.Println("")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	if f, err := Open(""); err == nil {
		fmt.Println(f)
	}

	for {
		break
	}

	user := map[string]interface{}{
		"username": "Jim Waters",
		"email":    "apietne@zam.fk",
		"country":  "Liberia",
	}

	for key, value := range user {
		fmt.Printf("%s => %v\n", key, value)
	}

	server := make(map[string]interface{})

	server["hostname"] = "localhost"
	server["port"] = 5713
	server["proxy"] = []string{
		"",
		"",
		"",
	}

	delete(server, "port")

	fmt.Println(server)

	var browser Browser

	browser.Name = "chrome"
	browser.Version = "161.11.0"
	browser.Engine = "v8"

	opera := Browser{
		Name:    "opera",
		Version: "100.45.11",
	}

	(&opera).Version = "89.21.67"

	ie := new(Browser)

	ie.Name = "IE"

	ie.Update("1.0.0")

	fmt.Println(browser.Engine == "", browser, opera, *ie)
	fmt.Println(ie.Debug)

	var comp Component = &Input{}

	fmt.Println(comp.Render(), Typeof(.0), Typeof(false), Isbool(' '), Isbool(true))

}
