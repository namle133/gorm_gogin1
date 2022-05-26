package database

import (
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (db *gorm.DB) {
	dsn := "host=localhost user=postgres password=Namle311 dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.Product{})
	return db
}
