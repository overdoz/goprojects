package main

import (
	"io"
	"net/http"
)



func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo dog dog")

	// fmt.Fprintf(w, "<h1>any Code & %s </h1>", r.Method)
}



func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat miau")
}

func main() {


	// dog route handles d function
	http.HandleFunc("/dog/", d)

	// car route handles c function
	http.HandleFunc("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
