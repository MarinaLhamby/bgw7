package storage

import (
	"encoding/json"
	"os"

	"github.com/bgw7/exercicio_2_falha/internal/domain"
	"github.com/bgw7/exercicio_2_falha/pkg/web/response"
)

type StorageJson struct {
}

func NewStorageJson() *StorageJson {
	return &StorageJson{}
}

func (s *StorageJson) Get() ([]domain.Product, error) {
	var products []domain.Product
	f, err := os.Open("../../internal/storage/products.json")

	if err != nil {
		return products, response.ErrFile
	}

	defer f.Close()
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		return products, response.ErrorDecoding.Format("products")
	}

	return products, nil
}

func (s *StorageJson) Save(products map[int64]domain.Product) error {
	var productsSlice []domain.Product

	for _, product := range products {
		productsSlice = append(productsSlice, product)
	}

	jsonString, err := json.Marshal(productsSlice)
	if err != nil {
		return response.ErrorEncoding.Format("json string")
	}
	os.WriteFile("../../internal/storage/products.json", jsonString, os.ModePerm)
	return nil
}
