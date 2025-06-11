package product

import (
	"strconv"

	"github.com/bgw7/exercicio_3_resposta/internal/domain"
	"github.com/bgw7/exercicio_3_resposta/internal/storage"
	"github.com/bgw7/exercicio_3_resposta/pkg/web/response"
)

type ProductJsonRepository struct {
	products           map[int64]domain.Product
	lastID             int64
	existingCodeValues map[string]bool
	jsonStorage        storage.Storage
}

func NewProductJsonRepository(storage storage.Storage) *ProductJsonRepository {
	repo := ProductJsonRepository{
		products:           make(map[int64]domain.Product),
		lastID:             0,
		existingCodeValues: make(map[string]bool),
		jsonStorage:        storage,
	}

	products, err := repo.jsonStorage.Get()
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

	err := pjr.jsonStorage.Save(pjr.products)
	if err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (pjr *ProductJsonRepository) UpdateProduct(product domain.Product) (domain.Product, error) {
	existingProduct := pjr.GetProductById(product.ID)
	if existingProduct == nil {
		return domain.Product{}, response.ErrEntityNotFound.Format("product", strconv.Itoa(int(product.ID)))
	}

	if existingProduct.CodeValue != product.CodeValue {
		if err := pjr.validateCodeValue(product.CodeValue); err != nil {
			return domain.Product{}, err
		}
		pjr.existingCodeValues[existingProduct.CodeValue] = false
		pjr.existingCodeValues[product.CodeValue] = true
	}

	pjr.products[product.ID] = product
	err := pjr.jsonStorage.Save(pjr.products)
	if err != nil {
		return domain.Product{}, err
	}

	productAfterUpdate := pjr.GetProductById(product.ID)
	return *productAfterUpdate, nil
}

func (pjr *ProductJsonRepository) PartialUpdateProduct(ID int64, partialProduct domain.PartialProduct) (domain.Product, error) {
	existingProduct := pjr.GetProductById(ID)
	if existingProduct == nil {
		return domain.Product{}, response.ErrEntityNotFound.Format("product", strconv.Itoa(int(ID)))
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

	err := pjr.jsonStorage.Save(pjr.products)
	if err != nil {
		return domain.Product{}, err
	}

	productAfterUpdate := pjr.GetProductById(ID)
	return *productAfterUpdate, nil
}

func (pjr *ProductJsonRepository) DeleteProduct(ID int64) error {
	existingProduct := pjr.GetProductById(ID)
	if existingProduct == nil {
		return response.ErrEntityNotFound.Format("product", strconv.Itoa(int(ID)))
	}

	delete(pjr.products, ID)

	err := pjr.jsonStorage.Save(pjr.products)
	if err != nil {
		return err
	}

	return nil
}

func (pjr *ProductJsonRepository) validateCodeValue(codeValue string) error {
	if pjr.existingCodeValues[codeValue] {
		return response.ErrProductCodeAlreadyExists
	}
	return nil
}
