package catalog

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100)" json:"name"`
	Active int8   `json:"active"` //sql-smallint

	Products     []Product     `gorm:"foreignKey:CategoryId"`
	SubCategorys []SubCategory `gorm:"foreignKey:CategoryId"`
}
