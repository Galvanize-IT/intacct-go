package intacct

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
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
