package intacct

import (
	"encoding/xml"
)

// TODO operator consts or type?

// TODO the dream of nesting is not yet alive
type Logical struct {
	Filters  []Expression `xml:"expression"`
	Operator string       `xml:"-"`
}

// TODO errors ignored
func (l Logical) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Make sure the start token is named 'filter'
	start.Name = xml.Name{Local: "filter"}
	e.EncodeToken(start)

	// If there are filters, inject an operator
	if l.Operator == "" {
		l.Operator = "and"
	}
	logical := xml.StartElement{
		Name: xml.Name{Local: "logical"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "logical_operator"}, Value: l.Operator},
		},
	}
	if len(l.Filters) > 1 {
		e.EncodeToken(logical)
	}
	for _, filter := range l.Filters {
		e.Encode(filter)
	}
	if len(l.Filters) > 1 {
		e.EncodeToken(logical.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func AllOf(expressions ...Expression) (logical Logical) {
	logical.Operator = "and"
	logical.Filters = expressions
	return
}

func AnyOf(expressions ...Expression) (logical Logical) {
	logical.Operator = "or"
	logical.Filters = expressions
	return
}

type Expression struct {
	XMLName  xml.Name `xml:"expression"`
	Field    string   `xml:"field"`
	Operator string   `xml:"operator"`
	Value    string   `xml:"value"`
}

func (ex Expression) IsEmpty() bool {
	return ex.Field == "" && ex.Operator == "" && ex.Value == ""
}

func NewExpression(field, operator, value string) Expression {
	// TODO xml.EscapeText?
	return Expression{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}

// TODO string or String() interface?
func Equal(field, value string) Expression {
	return NewExpression(field, "=", value)
}

func GreaterThan(field, value string) Expression {
	return NewExpression(field, ">", value)
}

func GTE(field, value string) Expression {
	return NewExpression(field, ">=", value)
}

func LessThan(field, value string) Expression {
	return NewExpression(field, "<", value)
}

func LTE(field, value string) Expression {
	return NewExpression(field, "<=", value)
}

// TODO list all possible fields?
const (
	CustomerID  Field = "customerid"
	InvoiceNo   Field = "invoiceno"
	DateCreated Field = "datecreated"
	DatePosted  Field = "dateposted"
	DateDue     Field = "datedue"
	DatePaid    Field = "datepaid"
	State       Field = "state"
)

type Field string

func (f Field) Equals(value string) Expression {
	return Equal(string(f), value)
}

func (f Field) GreaterThan(value string) Expression {
	return GreaterThan(string(f), value)
}

func (f Field) GTE(value string) Expression {
	return GTE(string(f), value)
}

func (f Field) LessThan(value string) Expression {
	return LessThan(string(f), value)
}

func (f Field) LTE(value string) Expression {
	return LTE(string(f), value)
}
