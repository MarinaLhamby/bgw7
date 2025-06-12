package product

import (
	"github.com/bgw7/exercicio_1_2_auth_response/internal/domain"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductById(id int64) *domain.Product
	FilterByPrice(minPrice float64) []domain.Product
	AddProduct(product domain.Product) (int64, error)
	UpdateProduct(product domain.Product) (domain.Product, error)
	PartialUpdateProduct(ID int64, partialProduct domain.PartialProduct) (domain.Product, error)
	DeleteProduct(ID int64) error
}
