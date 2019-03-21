package main

import (
	"encoding/xml"
	"github.com/prometheus/common/log"
	"os"
)

type VdpResponse struct {
	XMLName xml.Name `xml:"vdpResponse"`
	Header Header `cml:"header"`
}

type Header struct {
	Action string `xml:"action"`
}

const xmlFile = "./EES.xml"

func main() {
	h := Header{"Test"}
	v := VdpResponse{xml.Name{"ns1", "vdpResponse"}, h}



	file,err := os.Create(xmlFile)
	if err != nil {
		log.Fatalln(err)
	}

	xmlEnc := xml.NewEncoder(file)
	xmlEnc.Encode(v)
}
