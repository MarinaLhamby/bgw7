package internal

// ProductAttributes is the struct that represents the attributes of a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string `json:"description"`
	// Price is the price of the product.
	Price float64 `json:"price"`
}

// Product is the struct that represents a product.
type Product struct {
	// Id is the unique identifier of the product.
	Id int `json:"id"`
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

type BestSellers struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}
