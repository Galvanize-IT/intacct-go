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
type Method interface{}

// TODO Or Common method and change the Name field...?
type GetList struct {
	XMLName  xml.Name   `xml:"get_list"`
	Object   string     `xml:"object,attr"`
	MaxItems uint64     `xml:"maxitems,attr"`
	Filter   Expression `xml:"filter"` // TODO Are multiple allowed? nested?
	Sorts    Sorts      `xml:"sorts"`
}
