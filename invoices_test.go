package intacct

import (
	"encoding/xml"
	"testing"
)

var exampleInvoice = []byte(`<invoice>
	<key>38</key>
	<customerid>C-01</customerid>
	<datecreated>
		<year>2016</year>
		<month>03</month>
		<day>01</day>
	</datecreated>
	<dateposted>
		<year>2016</year>
		<month>03</month>
		<day>31</day>
	</dateposted>
	<datedue>
		<year>2016</year>
		<month>03</month>
		<day>01</day>
	</datedue>
	<datepaid>
		<year/>
		<month/>
		<day/>
	</datepaid>
	<termname>Due On Receipt</termname>
	<batchkey>4937</batchkey>
	<invoiceno>INV-02</invoiceno>
	<ponumber/>
	<totalamount>6400</totalamount>
	<totalpaid>0</totalpaid>
	<totaldue>6400</totaldue>
	<totalselected>0</totalselected>
	<description>Product</description>
	<trx_totalamount>6400</trx_totalamount>
	<trx_totalpaid>0</trx_totalpaid>
	<trx_totaldue>6400</trx_totaldue>
	<trx_totalselected>0</trx_totalselected>
	<basecurr>USD</basecurr>
	<currency>USD</currency>
	<billto>
		<contactname>Customer</contactname>
	</billto>
	<shipto>
		<contactname>Customer</contactname>
	</shipto>
	<whenmodified>02/23/2016 18:48:48</whenmodified>
	<state>Posted</state>
	<invoiceitems>
		<lineitem>
			<line_num>1</line_num>
			<accountlabel/>
			<glaccountno>40</glaccountno>
			<amount>3733.33</amount>
			<memo>Product 1</memo>
			<locationid>14</locationid>
			<departmentid>1</departmentid>
			<key>16</key>
			<totalpaid>0</totalpaid>
			<totaldue>3733.33</totaldue>
			<trx_amount>3733.33</trx_amount>
			<trx_totalpaid>0</trx_totalpaid>
			<trx_totaldue>3733.33</trx_totaldue>
			<currency>USD</currency>
			<projectkey/>
			<customerkey>12</customerkey>
			<vendorkey/>
			<employeekey/>
			<itemkey/>
			<classkey/>
			<billable/>
			<offsetglaccountno>131</offsetglaccountno>
		</lineitem>
		<lineitem>
			<line_num>2</line_num>
			<accountlabel/>
			<glaccountno>40</glaccountno>
			<amount>2666.67</amount>
			<memo>Product 2</memo>
			<locationid>14</locationid>
			<departmentid>1</departmentid>
			<key>17</key>
			<totalpaid>0</totalpaid>
			<totaldue>2666.67</totaldue>
			<trx_amount>2666.67</trx_amount>
			<trx_totalpaid>0</trx_totalpaid>
			<trx_totaldue>2666.67</trx_totaldue>
			<currency>USD</currency>
			<projectkey/>
			<customerkey>12</customerkey>
			<vendorkey/>
			<employeekey/>
			<itemkey/>
			<classkey/>
			<billable/>
			<offsetglaccountno>131</offsetglaccountno>
		</lineitem>
	</invoiceitems>
</invoice>`)

func TestInvoices(t *testing.T) {
	var invoice Invoice
	if err := xml.Unmarshal(exampleInvoice, &invoice); err != nil {
		t.Error(err)
	}

	if invoice.State != Posted {
		t.Errorf("unexpected invoice state: %s", invoice.State)
	}

	if len(invoice.Items) != 2 {
		t.Errorf(
			"unexpected number of items: %d != 2", len(invoice.Items),
		)
	}
}
