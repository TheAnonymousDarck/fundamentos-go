package database

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	ID            int `gorm:"primarykey; autoincrement"`
	IdProduct     string
	Price         float64
	BasePrice     float64
	OriginalPrice float64
	DateTime      string
}
