package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"net/http"
)

type Handler func() interface{}

type Koa struct {
	routes map[string]map[string]Handler
}

func (k *Koa) Get(url string, handler Handler) {
	if k.routes == nil {
		k.routes = make(map[string]map[string]Handler)
	}

	if k.routes[url] == nil {
		k.routes[url] = make(map[string]Handler)
	}

	k.routes[url]["GET"] = handler
}

func (k *Koa) Listen(host string, port int, callback func()) {
	mux := http.NewServeMux()

	for pattern, conf := range k.routes {
		func(patter string, conf map[string]Handler) {
			mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
				handler, ok := conf[r.Method]

				if ok {
					response := handler()

					switch response.(type) {
					case int:
					case string:
						fmt.Fprint(w, response)
					default:
						if json, ok := json.Marshal(response); ok == nil {
							fmt.Fprint(w, string(json))
						}
					}
				} else {
					w.WriteHeader(http.StatusNotImplemented)
					w.Write([]byte("method is not implemented"))
				}
			})
		}(pattern, conf)
	}

	mux.Handle("/404", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not found"))
	}))

	addr := fmt.Sprintf("%s:%d", host, port)

	callback()

	http.ListenAndServe(addr, mux)
}
