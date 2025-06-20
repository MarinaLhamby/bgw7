package product

import (
	"errors"
	"fmt"
	"regexp"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *Product) IsZero() bool {
	return p.ID == 0 &&
		p.Name == "" &&
		p.Quantity == 0 &&
		p.CodeValue == "" &&
		p.IsPublished == false &&
		p.Expiration == "" &&
		p.Price == 0.0
}

type PostRequest struct {
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (pr *PostRequest) Validate() error {
	if pr.Name == "" {
		return shouldBeInformedError("name")
	}

	if pr.Quantity == 0 {
		return shouldBeInformedError("quantity")
	}

	if pr.CodeValue == "" {
		return shouldBeInformedError("code_value")
	}

	dateRegex := regexp.MustCompile(`^\d{2}\/(0[1-9]|1[0-2])\/\d{4}$`)
	if pr.Expiration == "" {
		return shouldBeInformedError("expiration")
	}

	if !dateRegex.MatchString(pr.Expiration) {
		return errors.New("expiration should have the patter dd/mm/aaaa")
	}

	if pr.Price == 0.0 {
		return shouldBeInformedError("price")
	}

	return nil
}

func (pr *PostRequest) ToProduct() Product {
	return Product{
		Name:        pr.Name,
		Quantity:    pr.Quantity,
		CodeValue:   pr.CodeValue,
		IsPublished: pr.IsPublished,
		Expiration:  pr.Expiration,
		Price:       pr.Price,
	}
}

func shouldBeInformedError(field string) error {
	return errors.New(fmt.Sprintf("%s should be informed", field))
}

type PostResponse struct {
	ID int64 `json:"id"`
}
