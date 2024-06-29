package main

func main() {
	Printf("%v %v", keywords, defaultEncoding)
	Printf("%#v", Range(1, 10, 2))
	Printf("%T %v", Range, NewSet(0, false, -9i, 6.21))

	count, actions := CreateCounter(0)

	Printf("%v", count)

	actions["inc"]()

	Printf("%v", count)

}
