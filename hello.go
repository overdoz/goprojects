package main

import (
	"fmt"
	"math"

	"github.com/user/stringutil"
)

// Pointer Empfänger

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println("Hello, world.")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

}
