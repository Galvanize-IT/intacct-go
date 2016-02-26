package intacct

import (
	"encoding/xml"
)

type Request struct {
	XMLName   xml.Name `xml:"request"`
	Control   Control
	Operation RequestOperation
}

func NewRequestV2(config Config, fn interface{}) Request {
	login := NewLogin(config.User, config.Company, config.UserPassword)
	if config.Location != "" {
		login.LocationID = config.Location // TODO remove the ID suffixes?
	}
	return Request{
		Control: NewControlV2(config.Sender, config.SenderPassword),
		// TODO Multiple functions?
		Operation: RequestOperation{
			Authentication: Authentication{
				Login: login,
			},
			Content: Content{
				Functions: []interface{}{fn},
			},
		},
	}
}

type RequestOperation struct {
	XMLName        xml.Name `xml:"operation"`
	Authentication Authentication
	Content        Content `xml:"content"`
}

type Content struct {
	Functions []interface{} `xml:"function"`
}
