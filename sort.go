package intacct

import (
	"encoding/xml"
)

const (
	Asc  = "asc"
	Desc = "desc"
)

type SortField struct {
	XMLName xml.Name `xml:"sortfield"`
	Order   string   `xml:"order,attr"`
	Value   string   `xml:",chardata"`
}

// TODO Sort Asc and Desc could also be methods of fields...

type Sorts struct {
	Fields []SortField
}
