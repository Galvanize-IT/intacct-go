package intacct

import (
	"encoding/xml"
)

type Function struct {
	XMLName xml.Name `xml:"function"`
	Method
	ControlID string `xml:"controlid,attr"`
}

// TODO Method is not a good name...
// TODO Having an empty interface is also terrible
type Method interface{}
