package intacct

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

// TODO Case sensitivity?
type DateParts struct {
	Year  int        `xml:"year"`
	Day   int        `xml:"day"`
	Month time.Month `xml:"month"`
}

func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// TODO Separate into parts and encode?
	// TODO To capitalize or not?
	year := xml.StartElement{Name: xml.Name{Local: "year"}}
	month := xml.StartElement{Name: xml.Name{Local: "month"}}
	day := xml.StartElement{Name: xml.Name{Local: "day"}}

	// TODO errors ignored
	e.EncodeToken(start)
	e.EncodeToken(year)
	e.EncodeToken(xml.CharData([]byte(fmt.Sprintf("%d", date.Year()))))
	e.EncodeToken(year.End())
	e.EncodeToken(month)
	e.EncodeToken(xml.CharData([]byte(fmt.Sprintf("%d", date.Month()))))
	e.EncodeToken(month.End())
	e.EncodeToken(day)
	e.EncodeToken(xml.CharData([]byte(fmt.Sprintf("%d", date.Day()))))
	e.EncodeToken(day.End())
	e.EncodeToken(start.End())
	return nil
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var p DateParts
	if err := d.DecodeElement(&p, &start); err != nil {
		return err
	}
	date.Time = time.Date(p.Year, p.Month, p.Day, 0, 0, 0, 0, time.UTC)
	return nil
}
