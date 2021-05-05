package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo")
}

func barHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Bar")
}

func main() {
	mux := http.NewServeMux()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Fprint(rw, "Hello world")

	})

	http.HandleFunc("/bar", barHandler)

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
