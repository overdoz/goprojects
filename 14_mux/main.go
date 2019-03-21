package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo dog dog")

	// fmt.Fprintf(w, "<h1>any Code & %s </h1>", r.Method)
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat miau")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	// jede Route nach /dog/ führt zu dog
	mux.Handle("/dog/", d)

	// alles was an /cat angefügt wird sendet ein 404
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
