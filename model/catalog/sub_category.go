package catalog

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100)" json:"name"`
	Active int8   `json:"active"` //sql-smallint

	Products   []Product `gorm:"foreignKey:SubCategoryId"`
	CategoryId uint      `json:"category_id"`
}
