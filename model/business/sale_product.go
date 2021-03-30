package business

import "gorm.io/gorm"

type SaleProduct struct {
	gorm.Model
	SaleId   uint    `json:"sale_id"`
	Ticket   uint    `json:"ticket"`
	Code     string  `gorm:"type:varchar(20)" json:"code"`       //codigo barra
	Price    float32 `gorm:"type:decimal(12,2)" json:"price"`    //precio unitario
	Amount   float32 `gorm:"type:decimal(12,4)" json:"amount"`   //cantidad de producto
	Discount float32 `gorm:"type:decimal(12,2)" json:"discount"` //descuento en pesos
	Cost     float32 `gorm:"type:decimal(12,2)" json:"cost"`     //costo unitario
	Original float32 `gorm:"type:decimal(12,2)" json:"Original"` //precio unitario original
	Total    float32 `gorm:"type:decimal(12,2)" json:"total"`    //total venta o ticket
	Tax      float32 `gorm:"type:decimal(12,2)" json:"tax"`      //este es el iva
	NickName string  `gorm:"type:varchar(20)" json:"nick_name"`  //

}
