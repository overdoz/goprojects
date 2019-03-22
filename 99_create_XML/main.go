package main

import (
	// "fmt"
	"github.com/beevik/etree"
	// "io/ioutil"
	// "os"
)


func createXML(result string) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8" standalone="yes"`)

	vdpRes := doc.CreateElement("ns1:vdpResponse")
	header := vdpRes.CreateElement("ns1:header")

	action := header.CreateElement("ns1:action")
	action.CreateAttr("ns1:name", "DDA")

	subaction := action.CreateElement("ns1:subaction")
	subaction.CreateAttr("ns1:name", "RAR")

	eeSnapshot := header.CreateElement("ns1:eeSnapshot")
	eesUid := eeSnapshot.CreateElement("ns1:eesUid")
	eesUid.CreateText("WDDVP9AB1EJ001221;2014-10-24T12:29:49.650+0200")
	eesVersion := eeSnapshot.CreateElement("ns1:eesVersion")
	eesVersion.CreateText("2.2.0")


	trackingID := header.CreateElement("ns1:trackingID")
	trackingID.CreateText("WDDVP9AB1EJ001221_Tassi2")
	sourceID := header.CreateElement("ns1:trackingID")
	sourceID.CreateText("Test")

	responseState := vdpRes.CreateElement("ns1:responseState")
	responseState.CreateAttr("ns1:state", "0")
	messageItem := responseState.CreateElement("ns1:messageItem")
	messageSource := messageItem.CreateElement("ns1:messageSource")
	messageSource.CreateText("DDA")

	requestHeader := vdpRes.CreateElement("ns1:requestHeader")

	action2 := requestHeader.CreateElement("ns1:action")
	action2.CreateAttr("ns1:name", "DDA")
	subaction2 := action2.CreateElement("ns1:subaction")
	subaction2.CreateAttr("ns1:name", "RAR")


	eeSnapshot2 := requestHeader.CreateElement("ns1:eeSnapshot")
	eesUid2 := eeSnapshot2.CreateElement("ns1:eesUid")
	eesUid2.CreateText("WDDVP9AB1EJ001221")
	eesVersion2 := eeSnapshot2.CreateElement("ns1:eesVersion")
	eesVersion2.CreateText("2.2.0")


	trackingID2 := requestHeader.CreateElement("ns1:trackingID")
	trackingID2.CreateText("WDDVP9AB1EJ001221_Tassi2")
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

	/*	ees := etree.NewDocument()
		xmlFile, err := os.Open("./test.xml")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(xmlFile)

		if err := ees.ReadFromBytes(byteValue); err != nil {
			panic(err)
		}*/

	payload.CreateText(result)


	doc.WriteToFile("./EES2.xml")
}



func main() {

}
