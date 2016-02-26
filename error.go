package intacct

import "fmt"

type Error struct {
	Number       string `xml:"error>errorno"`
	Description  string `xml:"error>description"`
	Description2 string `xml:"error>description2"`
	Correction   string `xml:"error>correction"`
}

func (err Error) String() string {
	return fmt.Sprintf(
		"%s: %s %s - %s",
		err.Number,
		err.Description,
		err.Description2,
		err.Correction,
	)
}
