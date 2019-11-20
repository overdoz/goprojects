package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type server struct{}

//################## HTTP Handler ##############################################################################

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	case "POST":
		// w.WriteHeader(http.StatusCreated)
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
		text := r.FormValue("text")
		fmt.Println(text)
		printLKT(text)


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

func printLKT(t string) {
	printText := []byte(t)
	err := ioutil.WriteFile("./test.txt", printText, 0644)
	if err != nil {
		log.Fatal(err)
	}

	sh := `lp test.txt -d LKT`
	// sh := `echo "Hello World"`
	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", b)
}