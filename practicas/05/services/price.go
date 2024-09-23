package services

import (
	"gorm.io/gorm"
	"meliprices/database"
)

type PriceService interface {
	CreatePrice(price database.Price) error
}

type priceService struct {
	db *gorm.DB
}

func NewPriceService(db *gorm.DB) PriceService {
	return &priceService{db}
}

func (i *priceService) CreatePrice(price database.Price) error {
	err := i.db.Create(&price)

	return err.Error
}
