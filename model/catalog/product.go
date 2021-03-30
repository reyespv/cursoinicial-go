package catalog

import (
	bis "../../model/business"
	"gorm.io/gorm"
)

// Category model
type Product struct {
	gorm.Model
	Code          string  `gorm:"type:varchar(20)" json:"code"`
	Name          string  `gorm:"type:varchar(100)" json:"name"`
	Price         float32 `gorm:"type:decimal(12,2)" json:"price"`
	Cost          float32 `gorm:"type:decimal(12,2)" json:"cost"`
	Inventory     float32 `gorm:"type:decimal(12,2)" json:"inventory"`
	Active        int8    `json:"active"` //sql-smallint
	CategoryId    uint    `json:"category_id"`
	SubCategoryId uint    `json:"sub_category_id"`

	SaleProducts []*bis.SaleProduct `gorm:"many2many:product_sale_products;"`
}
