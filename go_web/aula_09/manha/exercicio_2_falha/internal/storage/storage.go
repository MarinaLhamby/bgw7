package storage

import "github.com/bgw7/exercicio_2_falha/internal/domain"

type Storage interface {
	Save(products map[int64]domain.Product) error
	Get() ([]domain.Product, error)
}
