package storage

import (
	"github.com/bgw7/exercicio_1_teste_sucesso/internal/domain"
)

type StorageMock struct {
}

func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

func (s *StorageMock) Get() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (s *StorageMock) Save(products map[int64]domain.Product) error {
	return nil
}
