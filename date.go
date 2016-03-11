package intacct

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

const americanDate = "1/2/2006"

// Use american format 1/2/2006
func (d Date) String() string {
	return d.Time.Format(americanDate)
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
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
		// TODO If parsing fails just leave the date zero
		return d.Skip()
	}
	date.Time = time.Date(p.Year, p.Month, p.Day, 0, 0, 0, 0, time.UTC)
	return nil
}
