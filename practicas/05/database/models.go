package database

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Id            string `gorm:"primarykey"`
	Price         float64
	BasePrice     float64
	OriginalPrice float64
	DateTime      string
}
