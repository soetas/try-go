package main

func main() {
	Printf("%s %s %T", "golang", Typeof(16_17_8_271), 16_17_8_271)

	var (
		count     int
		lng       float64
		completed bool
	)

	account, email := "Micheal Garza", "acfewhom@va.mh"

	Printf("%d %f %t", count, lng, completed)
	Printf("<%s, %s>", account, email)

	foregroundColor, backgroundColor := "black", "white"

	Swap(&foregroundColor, &backgroundColor)

	Printf("swap after: (%s, %s)", foregroundColor, backgroundColor)
	Printf("%v %v %v", keywords, Dracula, Typeof('a'))
	Printf("%v %c %s %v", 's', 's', `\n\t\0`, '\u0000')

	var gender byte = 'm'

	Printf("%s %s %s", defaultEncoding, Typeof(0i), Typeof(gender))

	for _, ch := range "GoWeb编程" {
		Printf("%c", ch)
	}

	// username := Scanf("username: ")
	// passwd := Scanf("passwd: ")
	// repos := Int(Scanf("repos of github: "))
	// commits := Float64(Scanf("commits of github: "))

	// Printf("(%s, %s) %v %T %v %T", username, passwd, repos, repos, commits, commits)

	var money Float = 198.2

	Printf("%v %T", money, money)

}
