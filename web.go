package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Pairs = url.Values

type HTTPRequest struct{}
type HTTPResponse struct{}

type Handler func(*HTTPRequest, *HTTPResponse) interface{}

type Database struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

type Koa struct {
	db     Database
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

func (k *Koa) Post(url string, handler Handler) {
	if k.routes == nil {
		k.routes = make(map[string]map[string]Handler)
	}

	if k.routes[url] == nil {
		k.routes[url] = make(map[string]Handler)
	}

	k.routes[url]["Post"] = handler
}

func (k *Koa) Listen(host string, port int, callback func()) {
	// initial database config
	db := k.db

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.Username, db.Password,
		db.Host, db.Port, db.Name)

	if db, err := sql.Open("mysql", datasource); err == nil {
		db.Ping()

		db.QueryRow("")

	}

	// multichannel route
	mux := http.NewServeMux()

	for pattern, conf := range k.routes {
		func(patter string, conf map[string]Handler) {
			mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
				handler, ok := conf[r.Method]

				if ok {
					req := &HTTPRequest{}
					res := &HTTPResponse{}

					response := handler(req, res)

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

type FetchOption struct {
	Method      string
	ContentType string
	Params      map[string]interface{}
	Body        map[string]interface{}
}

type FetchResponse struct {
	StatusCode int
	Status     string
	Header     http.Header
	Data       interface{}
}

type FileReader struct {
	io.Reader
	Total   int
	Size    int
	Current float64
}

func (f *FileReader) Read(s []byte) (n int, err error) {
	n, err = f.Reader.Read(s)

	f.Size += n
	f.Current = float64(f.Size) / float64(f.Total)

	fmt.Printf("current: %.2f%%\n", f.Current*100)

	return
}

func Stringify(value map[string]interface{}) string {
	pairs := make([]string, 0, len(value))

	for k, v := range value {
		pairs = append(pairs, fmt.Sprintf("%s=%v", k, v))
	}

	return strings.Join(pairs, "&")
}

func Fetch(url string, option FetchOption) FetchResponse {
	var res *http.Response
	var err error

	var (
		method      = option.Method
		params      = option.Params
		body        = option.Body
		contentType = option.ContentType
	)

	url += fmt.Sprintf("?%s", Stringify(params))

	switch strings.ToUpper(method) {
	case "GET":
		res, err = http.Get(url)
	case "POST":
		var str string

		if contentType == "application/x-www-form-encoding" {
			formData := make(Pairs)

			for k, v := range body {
				formData.Add(k, fmt.Sprintf("%v", v))
			}

			str = formData.Encode()
		}

		res, err = http.Post(url, contentType, strings.NewReader(str))

	case "PUT":
		req, _ := http.NewRequest(http.MethodPut, url, nil)

		req.Header.Add("Accept-Language", "zh-CN,en")

		res, err = http.DefaultClient.Do(req)
	default:
		panic("")
	}

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	return FetchResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Data:       string(content),
		Header:     res.Header,
	}
}

func Download(url, filename string) {
	res, _ := http.Get(url)

	file, _ := os.Create(filename)

	reader := FileReader{
		Reader: res.Body,
		Total:  int(res.ContentLength),
	}

	if _, err := io.Copy(file, &reader); err == nil {
		fmt.Printf("download finish ^_^ \n")
	}

}
