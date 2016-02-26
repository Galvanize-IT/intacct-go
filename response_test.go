package intacct

import (
	"encoding/xml"
	"testing"
)

var exampleData = []byte(`<data>
	<invoice>
		<key>5150</key>
	</invoice>
	<invoice>
		<key>5151</key>
	</invoice>
</data>`)

func TestData(t *testing.T) {
	var data Data
	if err := xml.Unmarshal(exampleData, &data); err != nil {
		t.Error(err)
	}
	if len(data.Invoices) != 2 {
		t.Error("expected 2 invoices in example data")
	}
}

var exampleResponse = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<response>
      <control>
            <status>success</status>
            <senderid>sender</senderid>
            <controlid>controlID</controlid>
            <uniqueid>false</uniqueid>
            <dtdversion>2.1</dtdversion>
      </control>
      <operation>
            <authentication>
                  <status>success</status>
                  <userid>user</userid>
                  <companyid>company</companyid>
                  <sessiontimestamp>2016-02-25T08:58:49-08:00</sessiontimestamp>
            </authentication>
            <result>
                  <status>success</status>
                  <function>get_list</function>
                  <controlid>testControlID</controlid>
                  <listtype start="0" end="0" total="1">invoice</listtype>
                  <data>
                        <invoice>
                              <key>5150</key>
                              <customerid>C-02</customerid>
                              <datecreated>
                                    <year>2013</year>
                                    <month>12</month>
                                    <day>01</day>
                              </datecreated>
                              <dateposted>
                                    <year>2014</year>
                                    <month>12</month>
                                    <day>01</day>
                              </dateposted>
                              <datedue>
                                    <year>2015</year>
                                    <month>12</month>
                                    <day>01</day>
                              </datedue>
                              <datepaid>
                                    <year>2016</year>
                                    <month>12</month>
                                    <day>01</day>
                              </datepaid>
                              <termname>Due On Receipt</termname>
                              <batchkey>16</batchkey>
                              <invoiceno>INV-02</invoiceno>
                              <ponumber>Monthly license fee</ponumber>
                              <totalamount>5000</totalamount>
                              <totalpaid>5000</totalpaid>
                              <totaldue>0</totaldue>
                              <totalselected>0</totalselected>
                              <description>Monthly license fee</description>
                              <trx_totalamount>5000</trx_totalamount>
                              <trx_totalpaid>5000</trx_totalpaid>
                              <trx_totaldue>0</trx_totaldue>
                              <trx_totalselected>0</trx_totalselected>
                              <basecurr>USD</basecurr>
                              <currency>USD</currency>
                              <billto>
                                    <contactname>Customer</contactname>
                              </billto>
                              <shipto>
                                    <contactname>Customer</contactname>
                              </shipto>
                              <whenmodified>01/14/2015 16:58:36</whenmodified>
                              <state>Paid</state>
                              <invoiceitems>
                                    <lineitem>
                                          <line_num>1</line_num>
                                          <accountlabel></accountlabel>
                                          <glaccountno>120</glaccountno>
                                          <amount>5000</amount>
                                          <memo>Product 001</memo>
                                          <locationid>10</locationid>
                                          <departmentid>11</departmentid>
                                          <key>13</key>
                                          <totalpaid>5000</totalpaid>
                                          <totaldue>0</totaldue>
                                          <trx_amount>5000</trx_amount>
                                          <trx_totalpaid>5000</trx_totalpaid>
                                          <trx_totaldue>0</trx_totaldue>
                                          <currency>USD</currency>
                                          <projectkey></projectkey>
                                          <customerkey>13</customerkey>
                                          <vendorkey></vendorkey>
                                          <employeekey></employeekey>
                                          <itemkey></itemkey>
                                          <classkey></classkey>
                                          <billable></billable>
                                          <offsetglaccountno>13</offsetglaccountno>
                                    </lineitem>
                              </invoiceitems>
                        </invoice>
                  </data>
            </result>
      </operation>
</response>`)

func TestResponse(t *testing.T) {
	var resp Response
	if err := xml.Unmarshal(exampleResponse, &resp); err != nil {
		t.Error(err)
	}

	if resp.Control.Status != Success {
		t.Errorf("unexpected control status: %s", resp.Control.Status)
	}

	if resp.Operation.Result.Status != Success {
		t.Errorf(
			"unexpected operation result status: %s",
			resp.Operation.Result.Status,
		)
	}

	if len(resp.Operation.Result.Data.Invoices) != 1 {
		t.Fatal("there should be 1 invoice in the result")
	}

	// TODO should probably be in TestInvoices
	invoice := resp.Operation.Result.Data.Invoices[0]
	if invoice.CustomerID != "C-02" {
		t.Errorf("unexpected customer: %s != C-02", invoice.CustomerID)
	}
	if invoice.InvoiceNumber != "INV-02" {
		t.Errorf("unexpected invoice: %s != INV-02", invoice.InvoiceNumber)
	}
	if invoice.TotalAmount != 5000.0 {
		t.Errorf("unexpected amount: %f != 5000", invoice.TotalAmount)
	}
	if invoice.DateCreated.Year() != 2013 {
		t.Errorf(
			"unexpected created year: %d != 2013",
			invoice.DateCreated.Year,
		)
	}
	if invoice.DatePosted.Year() != 2014 {
		t.Errorf(
			"unexpected post year: %d != 2014", invoice.DatePosted.Year,
		)
	}
	if invoice.DateDue.Year() != 2015 {
		t.Errorf(
			"unexpected due year: %d != 2015", invoice.DateDue.Year,
		)
	}
	if invoice.DatePaid.Year() != 2016 {
		t.Errorf(
			"unexpected paid year: %d != 2016", invoice.DatePaid.Year,
		)
	}
}

var exampleError = []byte(`<response>
  <control>
    <status>failure</status>
    <senderid>Galvanize</senderid>
    <controlid>controlID</controlid>
    <uniqueid>false</uniqueid>
    <dtdversion>2.1</dtdversion>
  </control>
  <errormessage>
    <error>
      <errorno>XL03000003</errorno>
      <description/>
      <description2>XML Parse error: Error 502: Value &amp;quot;whatever&amp;quot; for attribute object of get_list is not among the enumerated set. Line: 19, column: 70.</description2>
      <correction/>
    </error>
  </errormessage>
</response>`)

func TestResponseError(t *testing.T) {
	var resp Response
	if err := xml.Unmarshal(exampleError, &resp); err != nil {
		t.Error(err)
	}

	if resp.Control.Status != Failure {
		t.Errorf("unexpected control status: %s", resp.Control.Status)
	}

	if len(resp.Errors) != 1 {
		t.Fatal("there should be 1 error in the result")
	}
	err := resp.Errors[0]
	if err.Number != "XL03000003" {
		t.Errorf(
			"unexpected error number: %s != XL03000003", err.Number,
		)
	}
}
