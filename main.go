package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello,world"))
	})

	err := http.ListenAndServe("localhost:5713", mux)

	if err == nil {
		fmt.Printf("server is ready on http://localhost:5713/")
	} else {
		log.Fatal(err.Error())
	}

}
