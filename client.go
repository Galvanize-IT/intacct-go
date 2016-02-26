package intacct

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

const ContentType = `x-intacct-xml-request`

// TODO Does URL vary?
const apiURL = `https://api.intacct.com/ia/xml/xmlgw.phtml`

type Client struct {
	*http.Client
	config Config
	// TODO optional Backends
}

// NewRequest creates a request, but does not execute it
// TODO Errors?
// TODO accept method?
// TODO Pass operations instead?
func (c Client) NewRequest(m Method, ps ...Params) (*http.Request, error) {
	// Create the login
	login := NewLogin(c.config.User, c.config.Company, c.config.UserPassword)
	if c.config.Location != "" {
		login.LocationID = c.config.Location // TODO remove the ID suffixes?
	}

	// Create request body
	body := Request{
		Control: NewControlV2(c.config.Sender, c.config.SenderPassword),
		// TODO Multiple functions?
		Operation: Operation{
			Authentication: Authentication{
				Login: login,
			},
			Content: Content{
				Functions: []Function{
					{Method: m},
				},
			},
		},
	}

	b, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}

	// TODO Add buffer?
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", ContentType)
	return req, nil
}

func NewClient(config Config) Client {
	return Client{Client: &http.Client{}, config: config}
}

type API struct {
	Client
	Invoices  Invoices
	Customers Customers
}

func NewAPI(config Config) (api API) {
	// Pass the current client to each of the sub-clients
	client := NewClient(config)
	api.Client = client
	api.Invoices = Invoices{Client: client}
	api.Customers = Customers{Client: client}
	return api
}

// TODO Mock client
