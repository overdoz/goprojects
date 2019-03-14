package main

import (
	"fmt"
	"os"
	"text/template"
	"log"
)

func main() {
	var message string
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	tpl, err := template.ParseFiles("test.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(message)


}