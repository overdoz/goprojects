package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"os"
	"encoding/xml"

)

// Types ###########################################
//type Action struct {
//	XMLName xml.Name `xml: "action"`
//	SubAction string `xml: "subAction"`
//
//}
//
//type Subaction struct {
//	XMLName xml.Name `xml: "subAction"`
//
//
//}
//
//type VehicleID struct {
//	XMLName xml.Name `xml: "vehicleData"`
//	Type string `xml: "vin, attr"`
//}
type Users struct {
	Users xml.Name `xml:"users"`
	Userlist []User `xml:"person"`
}

type User struct {
	Name string `xml:"person"`
	Adresses Adress `xml:"adresses"`
}


type Adress struct {
	City string `xml:"city"`
	Street string `xml:"street"`
}

func Sum(end int) int {
	sum := 0
	for i := 0; i < end; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}


func main() {

	// xmlFile, err := os.Open("C:/Users/TL05566/go/src/goprojects/Hermes_ESR_Battery_VdpRequest.xml")
	xmlFile, err := os.Open("Users.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened xml file")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var user Users

	xml.Unmarshal(byteValue, &user)

	for i := 0; i < len(user.Userlist); i++ {
		fmt.Println(user.Userlist[i])
		// fmt.Println(user.Adresses.City)
	}









	// Stringers ####################################
	//hosts := map[string]IPAddr{
	//	"loopback":  {127, 0, 0, 1},
	//	"googleDNS": {8, 8, 8, 8},
	//}
	//for name, ip := range hosts {
	//	fmt.Printf("%v: %v\n", name, ip)
	//}
}

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

type IPAddr [4]byte

type MyFloat float64

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

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
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






