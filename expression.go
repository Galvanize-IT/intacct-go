package intacct

import ()

type Expression struct {
	Field    string `xml:"expression>field"`
	Operator string `xml:"expression>operator"`
	Value    string `xml:"expression>value"`
}

func (ex Expression) IsEmpty() bool {
	return ex.Field == "" && ex.Operator == "" && ex.Value == ""
}

func NewExpression(field, operator, value string) Expression {
	return Expression{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}

func Equal(field, value string) Expression {
	// xml.EscapeText?
	return NewExpression(field, "=", value)
}

// TODO list all possible fields?
const (
	CustomerID Field = "customerid"
	InvoiceNo  Field = "invoiceno"
)

type Field string

func (f Field) Equals(value string) Expression {
	return Equal(string(f), value)
}
