package main

import "fmt"

func main() {
	app := Koa{}

	app.Get("/", func() interface{} {
		return "hi, golang!"
	})

	app.Get("/login", func() interface{} {
		return "welcome to login"
	})

	app.Get("/api/users", func() interface{} {
		return []map[string]interface{}{
			{"username": "Lena May", "email": "zezam@rino.sk"},
			{"username": "Frederick Wood", "email": "naewu@panlowzeg.ru"},
			{"username": "Corey Steele", "email": "poj@vigkek.bw"},
		}
	})

	app.Listen("localhost", 5713, func() {
		fmt.Println("server is ready on http://localhost:5713/")
	})

}
