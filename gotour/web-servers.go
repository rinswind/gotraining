package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (this String) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, this)
}

func (this *Struct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, this)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
