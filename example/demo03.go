package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const host, port = "localhost", 5713

type IndexHandler struct{}
type AboutHandler struct{}

func (handler *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("")

	if err == nil {
		tpl.Execute(w, "")
	}

}

func (handler *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about me"))
}

type Operator interface {
	add(Number) Number
}

type Number int

func (n *Number) add(other Number) Number {
	return *n + other
}

func Main1() {
	addr := fmt.Sprintf("%s:%d", host, port)

	app := http.Server{
		Addr:    addr,
		Handler: nil,
	}

	http.Handle("/", &IndexHandler{})
	http.Handle("/about", &AboutHandler{})

	http.Handle("/download", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("start download ..."))
	}))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		protocol := r.URL.Scheme
		queryStr := r.URL.RawQuery
		query := r.URL.Query()
		acceptEncoding := r.Header["Accept-Encoding"]

		body := make([]byte, r.ContentLength)
		r.Body.Read(body)

		r.ParseForm()
		r.ParseMultipartForm(1024)

		fmt.Println(protocol, queryStr, query[""], acceptEncoding, string(body))

		http.Redirect(w, r, "", http.StatusMovedPermanently)

		w.Write([]byte("login"))

	})

	http.Handle("/404", http.NotFoundHandler())
	http.Handle("/redirect", http.RedirectHandler("", http.StatusMovedPermanently))
	http.Handle("/www", http.FileServer(http.Dir("")))

	err := app.ListenAndServe()

	if err != nil {
		panic("oh, something is error !")
	}

	title := "polyglot notebook vs jupyter"

	fmt.Println(title)
}
