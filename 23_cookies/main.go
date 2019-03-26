package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name: "Hello Cookie",
		Value: "Hello World",
	})
	fmt.Fprintln(res, "Not a virus")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")

}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Hello Cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(res, "Your cookie: ", c)

}