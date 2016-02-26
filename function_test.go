package intacct

import (
	"bytes"
	"encoding/xml"
	"testing"
)

var exampleFunction = `<function controlid="testControlID">
	<get_list object="invoice" maxitems="10">
		<filter>
			<expression>
				<field>customerid</field>
				<operator>=</operator>
				<value>C-01</value>
			</expression>
		</filter>
		<sorts>
			<sortfield order="desc">dateposted</sortfield>
		</sorts>
	</get_list>
</function>`

func TestFunction(t *testing.T) {
	getList := Function{
		ControlID: "testControlID",
		Method: GetList{
			Object:   "invoice",
			MaxItems: 10,
			Filter:   CustomerID.Equals("C-01"),
			Sorts:    Sorts{{Order: Desc, Value: "dateposted"}},
		},
	}

	b, err := xml.MarshalIndent(getList, "", "\t")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(b, []byte(exampleFunction)) {
		t.Errorf("Unexpected XML of function: %s", b)
	}
}
