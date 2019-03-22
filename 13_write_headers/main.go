package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Thanh", "Author of this page")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintf(w, "<h1>any Code & %s </h1>", r.Method)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
