# intacct-go
Go client for the Intacct API

*THIS IS TECHNOLOGY DEMO THAT CAN CHANGE AT ANY TIME*

Get all invoices for a customer:

```go
package main

import (
    "log"

    "github.com/Galvanize-IT/intacct-go"
)

func main() {
    config := intacct.Config{
        Sender:         "SENDER",
        SenderPassword: "SENDER_PASSWORD",
        User:           "xml_gateway",
        UserPassword:   "USER_PASSWORD",
        Company:        "COMPANY",
        Location:       "1000",
    }

    api := intacct.NewAPI(config)

    // List all invoices for a customer
    invoices, err := api.Invoices.List(
        intacct.ListParams{MaxItems: 100},
        intacct.CustomerID.Equals("C-01"),
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Returned %d invoices", len(invoices))
    for _, invoice := range invoices {
        log.Println(
            invoice.InvoiceNumber,
            invoice.DatePosted,
            invoice.State,
            invoice.TotalAmount,
            invoice.DatePaid,
        )
    }
}
```

Galvanize Software, 2016
