package product

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bgw7/exercicio_1_dominios/internal/domain"
)

var (
	repoProducts             = make(map[int64]domain.Product)
	lastID             int64 = 0
	existingCodeValues       = make(map[string]bool)
)

func LoadProducts() {
	f, err := os.Open("../../products.json")
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
		if product.ID > lastID {
			lastID = product.ID
		}
		repoProducts[product.ID] = product
		existingCodeValues[product.CodeValue] = true
	}

}

func GetAllProducts() []domain.Product {
	var products []domain.Product
	for _, product := range repoProducts {
		products = append(products, product)
	}
	return products
}

func GetProductById(id int64) *domain.Product {
	product := repoProducts[id]
	if product.IsZero() {
		return nil
	}
	return &product
}

func FilterByPrice(minPrice float64) []domain.Product {
	var filteredProducts []domain.Product

	for _, product := range repoProducts {
		if product.Price > minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func AddProduct(product domain.Product) (int64, error) {
	if existingCodeValues[product.CodeValue] {
		return 0, errors.New("code value already exists")
	}

	newID := lastID + 1

	product.ID = newID
	repoProducts[product.ID] = product
	existingCodeValues[product.CodeValue] = true
	lastID = product.ID

	return product.ID, nil
}
