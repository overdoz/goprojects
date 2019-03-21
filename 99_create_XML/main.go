package main

import (
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"os"
)



func main() {
	// el := etree.NewElement("ns1:")
	doc := etree.NewDocument()

	vdpRes := doc.CreateElement("ns1:vdpResponse")
	header := vdpRes.CreateElement("ns1:header")

	action := header.CreateElement("ns1:action")
		action.CreateAttr("ns1:name", "DDA")
		eeSnapshot := header.CreateElement("ns1:eeSnapshot")
			eeSnapshot.CreateText("Test")
		trackingID := header.CreateElement("ns1:trackingID")
			trackingID.CreateText("Test")
		sourceID := header.CreateElement("ns1:trackingID")
			sourceID.CreateText("Test")

	responseState := vdpRes.CreateElement("ns1:responseState")
		responseState.CreateText("Test")

	requestHeader := vdpRes.CreateElement("ns1:requestHeader")
		action2 := requestHeader.CreateElement("ns1:action")
		action2.CreateAttr("ns1:name", "DDA")
		eeSnapshot2 := requestHeader.CreateElement("ns1:eeSnapshot")
			eeSnapshot2.CreateText("Test")
		trackingID2 := requestHeader.CreateElement("ns1:trackingID")
			trackingID2.CreateText("Test")
		sourceID2 := requestHeader.CreateElement("ns1:trackingID")
			sourceID2.CreateText("Test")
		market := requestHeader.CreateElement("ns1:market")
			market.CreateText("Test")
		tenant := requestHeader.CreateElement("ns1:tenant")
			tenant.CreateText("Test")
		user := requestHeader.CreateElement("ns1:user")
			user.CreateText("Test")

	language := requestHeader.CreateElement("ns1:language")
		language.CreateText("Test")

	payload := vdpRes.CreateElement("ns1:payload")

	ees := etree.NewDocument()
	xmlFile, err := os.Open("./test.xml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(xmlFile)

	if err := ees.ReadFromBytes(byteValue); err != nil {
		panic(err)
	}

	payload.AddChild(ees)


	doc.WriteToFile("./EES2.xml")


}
