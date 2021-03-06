package main

import (
	"io"
	"net/http"
	"os"
)

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/test.jpg">`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("test.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/test.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}
