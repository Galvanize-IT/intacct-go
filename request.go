package intacct

import (
	"encoding/xml"
)

type Request struct {
	XMLName   xml.Name `xml:"request"`
	Control   Control
	Operation Operation
}

type Operation struct {
	XMLName        xml.Name `xml:"operation"`
	Authentication Authentication
	Content        Content `xml:"content"`
	Result         Result  `xml:"result,omitempty"`
}

type Content struct {
	Functions []Function
}
