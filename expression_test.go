package intacct

import (
	"bytes"
	"encoding/xml"
	"testing"
)

var exampleExpression = `<filter>
	<logical logical_operator="and">
		<expression>
			<field>dateposted</field>
			<operator>&gt;=</operator>
			<value>2/1/2016</value>
		</expression>
		<expression>
			<field>dateposted</field>
			<operator>&lt;</operator>
			<value>3/1/2016</value>
		</expression>
	</logical>
</filter>`

func TestExpression(t *testing.T) {
	exp := AllOf(
		DatePosted.GTE(NewDate(2016, 2, 1).String()),
		DatePosted.LessThan(NewDate(2016, 3, 1).String()),
	)

	b, err := xml.MarshalIndent(exp, "", "\t")
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(b, []byte(exampleExpression)) {
		t.Errorf("Unexpected XML of function: %s", b)
	}
}
