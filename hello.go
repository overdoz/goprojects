package main

import (
	"fmt"
)

type secretAgent struct {
	Name          string
	AllowedToKill bool
}

func (s *secretAgent) disallow() {
	s.AllowedToKill = false
}

func main() {
	x := 7
	// xi := []int{1,2,4,6,3}
	fmt.Printf("%T", x)

	bond := secretAgent{"James", true}
	fmt.Println(bond.AllowedToKill)
	bond.disallow()
	fmt.Println(bond.AllowedToKill)
}
