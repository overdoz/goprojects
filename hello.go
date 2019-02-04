package main

import (
	"fmt"
	// "github.com/user/exercise"
	"math"
)

// Types ###########################################

type Person struct {
	Name string
	Age  int
}

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type rect struct {
	width, height float64
}

type IPAddr [4]byte

type MyFloat float64

type ErrNegativeSqrt float64

type MyReader struct{}

// Pointer Empfänger ###############################

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Interface ######################################

// Was machen Interfaces genau?

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type geometry interface {
	area() float64
}

func (r rect) area() float64 {
	fmt.Println(r)
	return r.width * r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area)
}

// Interface-Werte: Interface führt die Methode mit einem konkreten Typen aus

// Type Switches

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Stringer ######################################

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Sprintf klappt mit jedem type
// a := Person{"Arthur Dent", 42}
// z := Person{"Zaphod Beeblebrox", 9001}
// fmt.Println(a, z)
// ------> Output: Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// Errors #######################################

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000000001 {
			return z, nil
		} else {
			z = z - (z*z-x)/(z*2)
		}
	}
}

// Reader ######################################

// Was macht der Reader? Warum nimmt er ein Byte Array?

// Übung:
func (r MyReader) Read(s []byte) (n int, err error) {
	s = s[:cap(s)]
	for i := range s {
		s[i] = 'A'
	}
	return cap(s), nil
}

func main() {

	r := rect{width: 10, height: 23}

	measure(r)

	do(21)
	do("Hello")
	do(true)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	// Stringers ####################################
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
