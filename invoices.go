package intacct

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	Paid   = "Paid"
	Posted = "Posted"
)

// TODO Are there multiple types of invoiceitems?
type InvoiceItem struct {
	LineItems []LineItem `xml:"lineitem"`
}

type LineItem struct {
	XMLName         xml.Name `xml:"lineitem"`
	LineNumber      string   `xml:"line_num"`
	AccountLabel    string   `xml:"accountlabel"`
	GLAccountNumber string   `xml:"glaccountno"`
	Amount          float64  `xml:"amount"`
	Memo            string   `xml:"memo"`
	LocationID      string   `xml:"locationid"`   // TODO int?
	DepartmentID    string   `xml:"departmentid"` // TODO int?
	Key             string   `xml:"key"`          // TODO int?
	TotalPaid       float64  `xml:"totalpaid"`
	TotalDue        float64  `xml:"totaldue"`
	// TRX?
	Currency    string `xml:"currency"`
	CustomerKey string `xml:"customerkey"`
}

type Invoice struct {
	XMLName       xml.Name `xml:"invoice"`
	Key           string   `xml:"key"`
	CustomerID    string   `xml:"customerid"`
	DateCreated   Date     `xml:"datecreated"`
	DatePosted    Date     `xml:"dateposted"`
	DateDue       Date     `xml:"datedue"`
	DatePaid      Date     `xml:"datepaid"`
	TermName      string   `xml:"termname"` // TODO ENUM?
	BatchKey      string   `xml:"batchkey"` // int?
	InvoiceNumber string   `xml:"invoiceno"`
	PONumber      string   `xml:"ponumber"`
	TotalAmount   float64  `xml:"totalamount"`
	TotalPaid     float64  `xml:"totalpaid"`
	TotalDue      float64  `xml:"totaldue"`
	Description   string   `xml:"description"`
	Currency      string   `xml:"currency"`
	State         string   `xml:"state"`
	// TODO modification date
}

type Invoices struct {
	Client
}

// Get returns an Invoice by invoice ID
func (inv Invoices) Get(id string) (Invoice, error) {
	// TODO We'll use the GetList command for now
	// TODO What about control IDs?
	get := Function{
		ControlID: "testControlID",
		Method: GetList{
			Object: "invoice",
			ListParams: ListParams{
				MaxItems: 2,
				Filter:   InvoiceNo.Equals(id),
			},
		},
	}

	// Create a new request using the Client
	req, err := inv.Client.NewRequest(get)
	if err != nil {
		return Invoice{}, err
	}

	resp, err := inv.Client.Do(req)
	if err != nil {
		return Invoice{}, err
	}
	defer resp.Body.Close()

	// TODO pull out status code and body status checks into client
	if resp.StatusCode != http.StatusOK {
		return Invoice{}, fmt.Errorf(
			"non-200 status code: %d", resp.StatusCode,
		)
	}

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Invoice{}, err
	}

	// Check the response for errors
	if err = inv.Client.CheckResponseErrors(body); err != nil {
		return Invoice{}, err
	}

	// Enforce one and only one result
	if len(body.Operation.Result.Data.Invoices) == 0 {
		return Invoice{}, fmt.Errorf(
			"no invoice was returned with the id %s", id,
		)
	} else if len(body.Operation.Result.Data.Invoices) > 1 {
		return Invoice{}, fmt.Errorf(
			"multiple invoices returned with the id %s", id,
		)
	}

	return body.Operation.Result.Data.Invoices[0], nil
}

// TODO Accept params - filtering and sorting
// Allow params:
// * ListParams
// * Expression
// * SortField
// TODO common list building
func (inv Invoices) List(params ...interface{}) ([]Invoice, error) {
	list := GetList{
		Object:     "invoice",
		ListParams: ListParams{MaxItems: 10}, // TODO Default page size?
	}

	for _, param := range params {
		switch p := param.(type) {
		case ListParams:
			list.ListParams = list.ListParams.Merge(p)
		case Expression:
			// TODO nested/multiple filters?
			list.ListParams.Filter = p
		case SortField:
			list.ListParams.Sorts = append(list.ListParams.Sorts, p)
		}
	}

	get := Function{
		ControlID: "testControlID",
		Method:    list,
	}

	// Create a new request using the Client
	req, err := inv.Client.NewRequest(get)
	if err != nil {
		return nil, err
	}

	resp, err := inv.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO pull out status code and body status checks into client
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"non-200 status code: %d", resp.StatusCode,
		)
	}

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	// Check the response for errors
	if err = inv.Client.CheckResponseErrors(body); err != nil {
		return nil, err
	}

	return body.Operation.Result.Data.Invoices, nil
}
