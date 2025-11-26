package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}
