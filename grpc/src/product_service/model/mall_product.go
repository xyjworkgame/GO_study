package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	CategoryId int64   `gorm:"column:category_id;type:int(11);not null" json:"category_id"`
	Name       string  `gorm:"column:name;type:varchar(100);not null;"json:"name"`
	Subtitle   string  `gorm:"column:subtitle;type:varchar(200)" json:"subtitle"`
	MainImage  string  `gorm:"column:main_image;type:varchar(500)" json:"main_image"`
	SubImages  string  `gorm:"column:sub_images;type:text" json:"sub_images"`
	Detail     string  `gorm:"column:detail;type:text" json:"detail"`
	Price      float64 `gorm:"column:price;type:decimal"json:"price"`
	Stock      int64   `gorm:"column:stock;type:int(11)" json:"stock"`
	Status     int     `gorm:"column:status;type:int(6)"json:"status"`
}
