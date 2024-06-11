package main

import "fmt"

type Widget struct {
	Display string
}

func (w *Widget) Paint() {
	if w == nil {
		return
	}
	fmt.Printf("%+v\n", w)
}

type Position struct {
	X float64
	Y float64
}

func (p *Position) SetX(x float64) {
	p.X = x
}

func (p *Position) SetY(y float64) {
	p.Y = y
}

type Button struct {
	Widget
	position Position
}

type Queriable interface {
	Query()
}

type Deletable interface {
	Delete()
}

type Updatable interface {
	Update()
}

type Insertable interface {
	Insert()
}

type Db interface {
	Queriable
	Deletable
	Updatable
	Insertable
}

type Mongodb struct{}

func (m *Mongodb) Query()  {}
func (m *Mongodb) Delete() {}
func (m *Mongodb) Update() {}
func (m *Mongodb) Insert() {}

type Sqlite struct{}

func (s *Sqlite) Query()  {}
func (s *Sqlite) Delete() {}
func (s *Sqlite) Update() {}
func (s *Sqlite) Insert() {}

type Mysql struct{}

func (m *Mysql) Query()  {}
func (m *Mysql) Delete() {}
func (m *Mysql) Update() {}
func (m *Mysql) Insert() {}

func Main10() {
	btn := Button{
		position: Position{0, 0},
		Widget: Widget{
			Display: "inline-block",
		},
	}

	btn.position.SetX(-1)

	fmt.Printf("%+v %v\n", btn, btn.Display)

	btn.Paint()
	btn.Widget.Paint()

	var db Db = &Mysql{}

	switch db.(type) {
	case *Mongodb:
		fmt.Printf("mongodb\n")
	case *Sqlite:
		fmt.Printf("sqlite\n")
	case *Mysql:
		fmt.Printf("mysql\n")
	default:
		fmt.Printf("")
	}

	var pint *int = new(int)
	var comparator func(x, y int) bool

	fmt.Printf("%T %#v %T %T\n", db, pint, nil, comparator)

}
