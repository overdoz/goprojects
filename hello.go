package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"os"
	"encoding/xml"

)

// Types ###########################################
type Action struct {
	XMLName xml.Name `xml: "v:action"`
	Type string `xml: "v:name,attr"`
	SubAction Subaction `xml: "v:subAction"`

}

type Subaction struct {
	XMLName xml.Name `xml: "v:subAction"`
	Type string `xml: "v:name,attr"`

}

type VehicleID struct {
	XMLName xml.Name `xml: "vdr:vehicleData"`
	Type string `xml: "vdr:vin, attr"`
}

type VehicleData struct {
	XMLName xml.Name `xml: ""`
	Acion string `xml: "v:action"`
	Subaction string `xml: "v:subAction"`
	VehicleID string
	PartNumber string

}

type Vehicle struct {
	XMLName xml.Name
	Acion string
	VehicleID string
	PartNumber string
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





func main() {

	xmlFile, err := os.Open("Hermes_ESR_Battery_VdpRequest.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully opened xml file")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var infos Action

	xml.Unmarshal(byteValue, &infos)

	fmt.Println(infos.Type)





	// Stringers #################################### 
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
