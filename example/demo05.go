package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

func Main01() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("template/index.html")

		template.Must(template.New("").Parse(""))

		tpl.ExecuteTemplate(w, "", "")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "template/login.html")

		case http.MethodPost:
			if err := r.ParseMultipartForm(1024); err == nil {
				formData := r.MultipartForm

				avatar := formData.File["avatar"][0]
				file, err := avatar.Open()

				if err == nil {
					data, _ := ioutil.ReadAll(file)
					fmt.Fprintln(w, string(data))
				}

				fmt.Fprintln(w, formData.Value, r.FormValue("account"), r.PostFormValue("passwd"))
			}

		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("method not implemented"))
		}
	})

	err := http.ListenAndServe("localhost:5713", nil)

	if err != nil {
		panic("Oh, error has happened!")
	}

}
