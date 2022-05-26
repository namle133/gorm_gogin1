package service

import (
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
)

type IProduct interface {
	Create(p *domain.Product) error
	Read() ([]domain.Product, error)
	ReadOne(id string) (*domain.Product, error)
	Delete(id string) error
	Update(id string, price int) error
}
