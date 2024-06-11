package designPattern

type widget interface {
	paint()
}

type button struct{}
type title struct{}
type input struct{}

func (b *button) paint() {}
func (t *title) paint()  {}
func (i *input) paint()  {}

func CreateWidget(name string) widget {
	switch name {
	case "button":
		return &button{}
	case "title":
		return &title{}
	case "input":
		return &input{}
	default:
		return nil
	}
}
