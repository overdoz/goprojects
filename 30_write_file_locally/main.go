package main

import (
	"fmt"
	"os"
)

func main() {
	
}

func writeFile(input string) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(input)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
