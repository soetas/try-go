package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"unsafe"
)

type any = interface{}
type object map[string]interface{}
type str string

func (s str) IsUpper() bool {
	for _, ch := range []rune(s) {
		if ch < 65 || ch > 90 {
			return false
		}
	}
	return true
}

func (s str) Concat(other str) str {
	return s + other
}

func (o object) HasKey(key string) bool {
	if _, exist := o[key]; exist {
		return true
	} else {
		return false
	}
}

type Slice struct {
	Start int
	End   int
	Step  int
}

func (s *Slice) String() string {
	return fmt.Sprintf("[%d:%d:%d]", s.Start, s.End, s.Step)
}

// factory function
func NewSlice(start, end, step int) *Slice {
	return &Slice{
		Start: start,
		End:   end,
		Step:  step,
	}
}

type OverflowError struct{}

func (e *OverflowError) Error() string {
	return "numerical overflow"
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// http server
type BottleHandler struct{}

func (b *BottleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type Bottle struct {
	app http.Server
}

func (b *Bottle) Get()  {}
func (b *Bottle) POst() {}
func (b *Bottle) Listen(callback func()) {
	b.app.ListenAndServe()
}

func NewBottle(host string, port int) *Bottle {
	return &Bottle{
		app: http.Server{
			Addr:    fmt.Sprintf("%s:%d", host, port),
			Handler: &BottleHandler{},
		},
	}
}

func Fetch(url, method string) {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		if res, err := http.Get(url); err == nil {
			defer res.Body.Close()

			content, _ := ioutil.ReadAll(res.Body)

			fmt.Println(string(content))
		}
	case "POST":

	}
}

func main() {
	fmt.Println("hi, soetas")

	var (
		account = "Susan Hudson"
		email   = "ra@fu.pa"
	)

	fmt.Printf("<%s, %s>\n", account, email)
	fmt.Println(unsafe.Sizeof(0), unsafe.Sizeof(1i))

	var firstName str = "Lenora"

	fmt.Println(firstName.Concat(" Cruz"), firstName.IsUpper())

	switch width := 980; {
	case width < 576:
		fmt.Println("x-small")
	case 576 <= width && width < 768:
		fmt.Println("small")
	case 768 <= width && width < 992:
		fmt.Println("medium")
	case 992 <= width && width < 1200:
		fmt.Println("large")
	case 1200 <= width && width < 1400:
		fmt.Println("extra large")
	case 1400 <= width:
		fmt.Println("extra extra large")
	}

	countries := make([]string, 5)

	countryList := append(countries, []string{"Guyana"}...)

	countryList[0] = "Luxembourg"

	fmt.Println(len(countries), cap(countries), countries)

	sprite := make(object)

	sprite["name"] = "Hamlet"

	fmt.Printf("%T %t\n", sprite, sprite.HasKey("name"))

	s := Slice{0, 10, 2}

	fmt.Println(s, s.String(), NewSlice(0, 10, 2))

	sort.Sort(sort.StringSlice{"", "", ""})

	var err error = &HTTPError{}

	if err, ok := err.(*HTTPError); ok {
		fmt.Println(err.Message)
	}

	app := NewBottle("localhost", 5713)

	app.Listen(func() {

	})

	Fetch("https://jsonplaceholder.typicode.com/todos/1", "get")

}
