package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bgw7/exercicio_2_storage/internal/domain"
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
		err := errors.New("error while opening file")
		return products, err
	}

	defer f.Close()
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		err := errors.New("error decoding products")
		return products, err
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
		return errors.New("error making the products as a json string")
	}
	os.WriteFile("../../internal/storage/products.json", jsonString, os.ModePerm)
	return nil
}
