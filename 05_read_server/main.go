package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// get Listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// loop over the listener and accept everything
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// handle the connection
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	// scanne die connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// scanne jedes Zeichen
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// fmt.Println("Code got here.")
}
