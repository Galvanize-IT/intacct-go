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

type Sorts []SortField

// Omit if zero length
func (sorts Sorts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(sorts) == 0 {
		return nil
	}
	// Otherwise unmarshal all items
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.Encode([]SortField(sorts)); err != nil {
		return err
	}
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}
	return nil
}

// TODO value type?
func SortAsc(value string) SortField {
	return SortField{Order: Asc, Value: value}
}

func SortDesc(value string) SortField {
	return SortField{Order: Desc, Value: value}
}
