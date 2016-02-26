package intacct

import (
	"bytes"
	"encoding/xml"
	"fmt"
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
	// Create request body
	body := NewRequestV2(c.config, m)

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

func (c Client) CheckResponseErrors(body Response) error {
	// Check the statuses
	// TODO Dump the body on non-success
	if body.Control.Status != Success {
		return fmt.Errorf(
			"unexpected control status (%s): %s",
			body.Control.Status, body.Errors,
		)
	}

	// TODO Where do all the errors hide?
	// if body.Operation.Result.Status != Success {
	// 	return fmt.Errorf(
	// 		"unexpected operation result status (%s): %s",
	// 		body.Operation.Result.Status, body.Errors,
	// 	)
	// }
	return nil
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
