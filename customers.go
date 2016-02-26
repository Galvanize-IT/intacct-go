package intacct

import ()

type Customer struct{}

type Customers struct { // TODO CustomersAPI?
	Client
}

// Get returns a Customer by customer ID
func (cust Customers) Get(id string) (Customer, error) {
	// TODO create a new request from the Client
	return Customer{}, nil
}

// TODO What about meta information? Attach to a List type?
func (cust Customers) List(params ...Params) ([]Customer, error) {
	return nil, nil
}
