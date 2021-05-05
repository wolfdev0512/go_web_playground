package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo")
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Fprint(rw, "Hello world")

	})

	http.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprint(rw, "Hello Bar")

	})
	http.ListenAndServe(":3000", nil)
}
