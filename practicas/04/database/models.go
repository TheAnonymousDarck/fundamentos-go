package database

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id string  `gorm:"primarykey"`
	Title      string 
	Price      float64
	Condition  string 
	CategoryId string 
	Url        string 
}