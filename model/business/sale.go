package business

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	Ticket   uint    `json:"ticket"`
	Client   uint    `json:"client"`
	Credit   string  `gorm:"type:varchar(1)"    json:"credit"`    //S es venta a crpedito N no es a credito
	Total    float32 `gorm:"type:decimal(12,2)" json:"total"`     //total venta o ticket
	Tax      float32 `gorm:"type:decimal(12,2)" json:"tax"`       //este es el iva
	Status   string  `gorm:"type:varchar(1)"    json:"status"`    //sql-smallint
	NickName string  `gorm:"type:varchar(20)"   json:"nick_name"` //sql-smallint
	Session  int16   `json:"session"`                             //Numero de corte

	SaleProducts []SaleProduct `gorm:"foreignKey:SaleId"`
}
