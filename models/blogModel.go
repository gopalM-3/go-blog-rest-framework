package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Category string
	Title 	 string
	Excerpt  string `gorm:"type:text"`
	Content  string
	Slug 	 string `gorm:"unique"`
	Author 	 string `gorm:"foreignKey:Username"`
	Hashtags string `gorm:"type:text"`
	Status 	 bool
	Votes  int
}