package main

import (
	"fmt"
	"os"
)

func main() {
	url := "https://httpbin.org/put"

	fmt.Printf("%+v\n", Fetch(url, FetchOption{
		Method: "put",
		Params: map[string]interface{}{
			"skip":  1,
			"limit": 3,
		},
	}).Header.Get("Content-Type"))

	Download("https://haowallpaper.com/link/common/file/getCroppingImg/15024072666287424", "poster.png")

	os.Exit(0)

	app := Koa{
		db: Database{
			Username: "root",
			Password: "123456+",
			Host:     "localhost",
			Port:     3306,
			Name:     "test",
		},
	}

	app.Get("/", func(req *HTTPRequest, res *HTTPResponse) interface{} {
		return "hi, golang!"
	})

	app.Get("/login", func(req *HTTPRequest, res *HTTPResponse) interface{} {
		return "welcome to login"
	})

	app.Post("/login", func(h1 *HTTPRequest, h2 *HTTPResponse) interface{} {
		return map[string]interface{}{
			"statusCode": 200,
			"message":    "success",
		}
	})

	app.Get("/api/users", func(req *HTTPRequest, res *HTTPResponse) interface{} {
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
