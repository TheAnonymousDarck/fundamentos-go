package services

import (
	"meliitems/database"
)

type ItemService interface {
	CreateItem(item database.Item) error
}

type itemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService(db)
}


func (i *itemService) CreateItem(item database,Item) error {
	err := i.db.Create(&item)

	return err.Error
}
