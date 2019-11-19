package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

//################## HTTP Handler ##############################################################################

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		//w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		//w.Write([]byte(`{"message": "get called"}`))
		http.ServeFile(w, r, "form.html")
	case "POST":
		w.WriteHeader(http.StatusCreated)
		// w.Write([]byte(`{"message": "post called"}`))
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
		fmt.Println(name)
		fmt.Println(address)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
