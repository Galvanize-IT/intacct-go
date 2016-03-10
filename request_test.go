package intacct

import (
	"encoding/xml"
	"testing"
)

var exampleRequest = `<request>
<control>
	<senderid></senderid>
	<password></password>
	<controlid>controlID</controlid>
	<uniqueid>false</uniqueid>
	<dtdversion>2.1</dtdversion>
</control>
<operation>
	<authentication>
		<login>
			<userid></userid>
			<companyid></companyid>
			<password></password>
		</login>
		<sessiontimestamp>0001-01-01T00:00:00Z</sessiontimestamp>
	</authentication>
	<content>
		<function controlid="testControlID">
			<get_list object="invoice" maxitems="2">
				<filter>
					<expression>
						<field>invoiceno</field>
						<operator>=</operator>
						<value>INV-01</value>
					</expression>
				</filter>
			</get_list>
		</function>
	</content>
</operation>
</request>`

func TestRequest(t *testing.T) {
	empty := Config{}

	get := Function{
		ControlID: "testControlID",
		Method: GetList{
			Object: "invoice",
			ListParams: ListParams{
				MaxItems: 2,
				Filter:   InvoiceNo.Equals("INV-01"),
			},
		},
	}

	body := NewRequestV2(empty, get)
	_, err := xml.MarshalIndent(body, "", "\t")
	if err != nil {
		t.Error(err)
	}
	// TODO test output
}
