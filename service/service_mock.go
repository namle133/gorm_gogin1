package service

import (
	"fmt"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
	"gorm.io/gorm"
)

type Product struct {
	Db *gorm.DB
}

func (p *Product) ReadOne(id string) (*domain.Product, error) {
	var item *domain.Product
	err := p.Db.First(&item, "code = ?", id).Error
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (p *Product) Create(item *domain.Product) error {
	err := p.Db.Create(item).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (p *Product) Read() ([]domain.Product, error) {
	var products []domain.Product
	err := p.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) Delete(id string) error {
	var item *domain.Product
	err := p.Db.Where("code = ?", id).Delete(&item).Error
	if err != nil {
		return err
	}
	return nil

}

func (p *Product) Update(id string, price int) error {
	var item *domain.Product
	err := p.Db.Model(&item).Where("code = ?", id).Update("price", price).Error
	fmt.Println(err)
	fmt.Println(id)
	fmt.Println(price)
	if err != nil {
		return err
	}
	return nil
}
