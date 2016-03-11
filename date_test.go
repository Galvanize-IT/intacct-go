package intacct

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	jan1 := NewDate(2012, 1, 1)
	if jan1.String() != "1/1/2012" {
		t.Errorf("Unexpected date string: %s != 1/1/2012", jan1)
	}
	dec12 := NewDate(2006, 12, 12)
	if dec12.String() != "12/12/2006" {
		t.Errorf("Unexpected date string: %s != 12/12/2006", dec12)
	}

	now := Date{time.Now()}
	b, err := xml.MarshalIndent(now, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	var parsed Date
	if err = xml.Unmarshal(b, &parsed); err != nil {
		t.Fatal(err)
	}

	if parsed.Year() != now.Year() {
		t.Errorf("unexpected year: %d != %d", parsed.Year(), now.Year())
	}
	if parsed.Month() != now.Month() {
		t.Errorf("unexpected month: %d != %d", parsed.Month(), now.Month())
	}
	if parsed.Day() != now.Day() {
		t.Errorf("unexpected day: %d != %d", parsed.Day(), now.Day())
	}
}
