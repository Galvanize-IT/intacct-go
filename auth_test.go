package intacct

import (
	"bytes"
	"encoding/xml"
	"testing"
)

var exampleLogin = `<login>
	<userid>test</userid>
	<companyid>company</companyid>
	<password>secret</password>
</login>`

var exampleLoginWithLocation = `<login>
	<userid>test</userid>
	<companyid>company</companyid>
	<password>secret</password>
	<locationid>1</locationid>
</login>`

func TestLogin(t *testing.T) {
	login := NewLogin("test", "company", "secret")
	b, err := xml.MarshalIndent(login, "", "\t")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(b, []byte(exampleLogin)) {
		t.Errorf("Unexpected XML of login: %s", b)
	}

	login.LocationID = "1"
	b, err = xml.MarshalIndent(login, "", "\t")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(b, []byte(exampleLoginWithLocation)) {
		t.Errorf("Unexpected XML of login with login: %s", b)
	}
}

var exampleControl = `<control>
	<senderid>sender</senderid>
	<password>secret</password>
	<controlid>controlID</controlid>
	<uniqueid>false</uniqueid>
	<dtdversion>2.1</dtdversion>
</control>`

func TestControl(t *testing.T) {
	control := NewControlV2("sender", "secret")
	b, err := xml.MarshalIndent(control, "", "\t")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(b, []byte(exampleControl)) {
		t.Errorf("Unexpected XML of control: %s", b)
	}
}
