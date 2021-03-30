package model

import (
	bis "../model/business"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName  string `gorm:"type:varchar(20);uniqueIndex"  json:"nick_name"`
	Email     string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	FirstName string `gorm:"size:100;not null"             json:"first_name"`
	LastName  string `gorm:"size:100;not null"             json:"last_name"`
	Password  string `gorm:"size:100;not null" 	           json:"password"`

	Sales []bis.Sale `gorm:"foreignKey:NickName;references:NickName"`
}
