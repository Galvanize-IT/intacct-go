package intacct

import (
	"encoding/xml"
)

// Status consts
const (
	Success = "success"
)

// TODO Set result type method?
type Response struct {
	XMLName   xml.Name `xml:"response"`
	Control   Control
	Operation Operation
}

// TODO Or use delayed parsing?
type Result struct {
	Status    string `xml:"status"`
	Function  string `xml:"function"`
	ControlID string `xml:"controlid"`
	// TODO listtype?
	Data Data `xml:"data"`
}

type Data struct {
	Invoices  []Invoice  `xml:"invoice"`
	Customers []Customer `xml:"customer"`
}
