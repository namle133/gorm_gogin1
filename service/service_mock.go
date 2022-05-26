package service

import (
	"fmt"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/database"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
	"gorm.io/gorm"
)

type product struct {
	db *gorm.DB
}

func Connect(db *gorm.DB) *product {
	p := &product{
		db: database.ConnectDatabase(),
	}
	return p
}

func (p *product) Create(item *domain.Product) error {
	fmt.Println("sfsdf")
	//p = &product{
	//	db: database.ConnectDatabase(),
	//}
	//fmt.Println(p)
	p.db = database.ConnectDatabase()
	err := p.db.Create(item).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (p *product) Read() ([]domain.Product, error) {
	var products []domain.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *product) Delete(id string) error {
	var item *domain.Product
	err := p.db.Where("code = ?", id).Delete(&item).Error
	if err != nil {
		return err
	}
	return nil

}

func (p *product) Update(id string, price int) error {
	var item *domain.Product
	err := p.db.Model(&item).Where("code = ?", id).Update("price", price).Error
	fmt.Println(err)
	fmt.Println(id)
	fmt.Println(price)
	if err != nil {
		return err
	}
	return nil
}
