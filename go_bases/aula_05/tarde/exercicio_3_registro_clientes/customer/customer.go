package customer

import "errors"

var ErrCustomerIsZero = errors.New("The customer is empty, please provide valid data")

type Customer struct {
	File        string
	Name        string
	ID          int64
	PhoneNumber string
	Address     string
}

func (c *Customer) Compare(customer Customer) bool {
	return c.Address == customer.Address &&
		c.File == customer.File &&
		c.ID == customer.ID &&
		c.PhoneNumber == customer.PhoneNumber &&
		c.Name == customer.Name
}

func (c *Customer) IsZero() error {
	if c.File == "" &&
		c.Name == "" &&
		c.ID == 0 &&
		c.PhoneNumber == "" &&
		c.Address == "" {
		return ErrCustomerIsZero
	}
	return nil
}
