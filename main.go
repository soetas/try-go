package main

import (
	"fmt"
	_ "stdlib/random"
)

const USER_AGENT = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 
(KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36`

const (
	MALE = iota
	FEMALE
)

var version = "1.16.6"

func divmod(x, y int) (int, int) {
	return x / y, x % y
}

func exit(code int) {
	fmt.Printf("exit code: %d\n", code)
}

func setButtonSize(button map[string]interface{}, size string) {
	button["size"] = size
}

type Creator struct {
	Name string
	Post
}

type Post struct {
	Id    int
	Title string
	Body  string
}

func (post *Post) setTitle(title string) {
	post.Title = title
}

type Any interface{}

func main() {
	defer exit(0)

	var count int
	var username, passwd string = "Danny Porter", "0cdeb7ce2a59"

	email := "socercep@wisuncuh.la"

	pint := &count

	*pint = 45

	fmt.Println(username, passwd, *pint, pint)
	fmt.Printf("%v %T %T %s\n", count, count, email, version)
	fmt.Println(USER_AGENT)

	fmt.Println(divmod(10, 4))

	scores := [...]float64{3: 98.1, 5: 78.0}

	for _, score := range scores {
		fmt.Println(score)
	}

	fmt.Printf("%T %T\n", scores, scores[:])

	colors := make([]string, 3, 5)

	colors = append(colors, "lightblue", "pink")

	colors[3:][0] = "blue"

	fmt.Printf("%v %d %d\n", colors, len(colors), cap(colors))

	player := make(map[string]string)

	player["level"] = "白金"

	button := map[string]interface{}{
		"size":     "large",
		"icon":     "spin",
		"disabled": false,
	}

	setButtonSize(button, "default")

	fmt.Println(player, button)

	post := Post{
		Id:    0,
		Title: "",
		Body:  "",
	}

	fmt.Println(post)

	post.setTitle("aut provident")

	fmt.Println(post)

	creator := Creator{"", Post{}}

	creator.setTitle("")

}

func init() {
	fmt.Println("main::init called")
}
