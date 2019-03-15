package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var message string

	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, message)
}
