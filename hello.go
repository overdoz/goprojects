package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"os"
	"encoding/xml"
	"strings"

)

// Types ###########################################
type VDPrequest struct {
	Head Header `xml:"header"`
	InnerPayload Payload `xml:"payload"`
}

type Action struct {
	InnerSub Subaction `xml:"subAction"`
	Name string `xml:"name,attr"`
}

type Header struct {
	InnerAction Action `xml:"action"`
}

type Subaction struct {
	Name string `xml:"name,attr"`
}

type Payload struct {
	Data VehicleData `xml:"vehicleData"`
}

type VehicleData struct {
	VehicleID string `xml:"vin,attr"`
}



type Users struct {
	//Users xml.Name `xml:"user"`
	//Userlist []User `xml:"user"`
	InnerXML string `xml:",innerxml"`
}

//type User struct {
//	Name string `xml:"person"`
//	Adresses Adress `xml:"adresses"`
//}


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

	xmlFile, err := os.Open("C:/Users/TL05566/go/src/goprojects/Hermes_ESR_Battery_VdpRequest.xml")
	// xmlFile, err := os.Open("Users.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened xml file")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	// var user Users
	var request VDPrequest

	xml.Unmarshal(byteValue, &request)
	fmt.Println(request.Head.InnerAction.InnerSub.Name)
	fmt.Println(request.Head.InnerAction.Name)
	fmt.Println(request.InnerPayload.Data.VehicleID)


	var input float64
	slice1 := []float64{}



	for len(slice1) < 10 {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%f", &input)
		slice1 = append(slice1, input)
	}
	fmt.Println(slice1)

	v:= make(map[string]int)
	text := "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"
	textArray := strings.Fields(text)

	for _, word := range textArray {
		v[word]++
	}

	for key, val := range v {
		fmt.Printf("key[%s] value: %v \n", key, val)
	}

	// fmt.Println()










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









