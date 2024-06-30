package main

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Address struct {
	Country string
	City    string
	Street  string
}

type User struct {
	*Address
	City string
}

func (u *User) program() {}

// func program() {}

func main() {
	src := []string{"python", "c++", "java"}
	dst := []string{"c#", "dart", "c", "javascript", "go"}

	customer := map[string]any{
		"name":   "Lola Chapman",
		"gender": "male",
		"products": []string{
			"short sleeve",
			"sweater",
			"hat",
		},
	}

	posts := [...]Post{
		3: {},
		5: {},
	}

	var pointerOfInt *int

	pointerOfArray := new([3]int)

	arrOfPointers := [...]*int{
		3: new(int),
	}

	copy(dst, src)

	Printf("%v %d %t", dst, len(customer), Hashable([]int{}))
	Printf("%t", HasKey(customer, "gender"))
	Printf("%v %t", posts, Post{Id: 0, UserId: 0, Title: "", Body: ""} == Post{})
	Printf("%p %t %T %T", &posts, IsNil(pointerOfInt), pointerOfArray, arrOfPointers)

	user := User{
		Address: &Address{
			Country: "China",
			City:    "Tj",
			Street:  "",
		},
		City: "Bj",
	}

	Printf("%v %v", user.Address.Country, user.City)

	user.program()

}
