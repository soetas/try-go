package main

import "strings"

type Model struct{}

type Post struct {
	Model
	Id    int
	Title string
	Body  string
}

func (m *Model) Save()  {}
func (m Model) Delete() {}

func (p *Post) Save() {
	Printf("save post record to database")
}

type Painter interface {
	Paint()
}

type Renderer interface {
	Render()
}

type Shaper interface {
	Painter
}

type Circle struct {
	X, Y   float64
	Radius float64
}

type Rect struct {
	X, Y   float64
	Width  float64
	Height float64
}

type Square struct {
	X, Y float64
	Side float64
}

func (c *Circle) Paint() {
	Printf("paint circle at (%.2f, %.2f)", c.X, c.Y)
}

func (r *Rect) Paint() {
	Printf("paint rect at (%.2f, %.2f)", r.X, r.Y)
}

func (t *Square) Paint() {
	Printf("paint square at (%.2f, %.2f)", t.X, t.Y)
}

func main() {
	post := Post{
		Model: Model{},
		Id:    1,
		Title: "",
		Body:  "",
	}

	post.Save()
	post.Model.Save()

	Printf("%t %T %T %T", Model{} == Model{}, post.Save, Model.Delete, (*Model).Save)

	var shape Shaper = &Circle{
		X: 56.18212,
		Y: 70.12178,
	}

	shape.Paint()

	Printf("%t %t", IsBool(false), IsBool(0))
	Printf("%t\n%s", strings.Contains("hi,golang", "go"), strings.Join([]string{}, " "))

}
