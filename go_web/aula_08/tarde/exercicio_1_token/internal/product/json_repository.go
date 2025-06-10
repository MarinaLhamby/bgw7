package product

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bgw7/exercicio_1_token/internal/domain"
)

type ProductJsonRepository struct {
	products           map[int64]domain.Product
	lastID             int64
	existingCodeValues map[string]bool
}

func NewProductJsonRepository() *ProductJsonRepository {
	repo := ProductJsonRepository{
		products:           make(map[int64]domain.Product),
		lastID:             0,
		existingCodeValues: make(map[string]bool),
	}

	f, err := os.Open("../../internal/product/products.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	var products []domain.Product
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		if product.ID > repo.lastID {
			repo.lastID = product.ID
		}
		repo.products[product.ID] = product
		repo.existingCodeValues[product.CodeValue] = true
	}
	return &repo
}

func (pjr *ProductJsonRepository) GetAllProducts() []domain.Product {
	var products []domain.Product
	for _, product := range pjr.products {
		products = append(products, product)
	}
	return products
}

func (pjr *ProductJsonRepository) GetProductById(id int64) *domain.Product {
	product, ok := pjr.products[id]
	if !ok {
		return nil
	}
	return &product
}

func (pjr *ProductJsonRepository) FilterByPrice(minPrice float64) []domain.Product {
	var filteredProducts []domain.Product

	for _, product := range pjr.products {
		if product.Price > minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (pjr *ProductJsonRepository) AddProduct(product domain.Product) (int64, error) {
	if err := pjr.validateCodeValue(product.CodeValue); err != nil {
		return 0, err
	}

	newID := pjr.lastID + 1

	product.ID = newID
	pjr.products[product.ID] = product
	pjr.existingCodeValues[product.CodeValue] = true
	pjr.lastID = product.ID

	return product.ID, nil
}

func (pjr *ProductJsonRepository) UpdateProduct(product domain.Product) (domain.Product, error) {
	existingProduct := pjr.GetProductById(product.ID)
	if existingProduct == nil {
		return domain.Product{}, errors.New("product doesn`t exist")
	}

	if existingProduct.CodeValue != product.CodeValue {
		if err := pjr.validateCodeValue(product.CodeValue); err != nil {
			return domain.Product{}, err
		}
		pjr.existingCodeValues[existingProduct.CodeValue] = false
		pjr.existingCodeValues[product.CodeValue] = true
	}

	pjr.products[product.ID] = product

	productAfterUpdate := pjr.GetProductById(product.ID)
	return *productAfterUpdate, nil
}

func (pjr *ProductJsonRepository) PartialUpdateProduct(ID int64, partialProduct domain.PartialProduct) (domain.Product, error) {
	existingProduct := pjr.GetProductById(ID)
	if existingProduct == nil {
		return domain.Product{}, errors.New("product doesn`t exist")
	}

	if partialProduct.CodeValue != "" {
		existingProduct.CodeValue = partialProduct.CodeValue
	}

	if partialProduct.Expiration != "" {
		existingProduct.Expiration = partialProduct.Expiration
	}

	if partialProduct.IsPublished != nil {
		existingProduct.IsPublished = *partialProduct.IsPublished
	}

	if partialProduct.Name != "" {
		existingProduct.Name = partialProduct.Name
	}

	if partialProduct.Price != nil {
		existingProduct.Price = *partialProduct.Price
	}

	if partialProduct.Quantity != nil {
		existingProduct.Quantity = *partialProduct.Quantity
	}

	pjr.UpdateProduct(*existingProduct)

	productAfterUpdate := pjr.GetProductById(ID)
	return *productAfterUpdate, nil
}

func (pjr *ProductJsonRepository) DeleteProduct(ID int64) error {
	existingProduct := pjr.GetProductById(ID)
	if existingProduct == nil {
		return errors.New("product doesn`t exist")
	}

	delete(pjr.products, ID)

	return nil
}

func (pjr *ProductJsonRepository) validateCodeValue(codeValue string) error {
	if pjr.existingCodeValues[codeValue] {
		return errors.New("code value already exists")
	}
	return nil
}
