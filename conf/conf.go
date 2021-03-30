package conf

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	mod "../model"
	bis "../model/business"
	cat "../model/catalog"
)

func ConnectDB(DbHost string, DbName string, DbUser string, DbPassword string, DbPort string) (*gorm.DB, error) {
	//mostramos en pantalla los datos de conexión
	fmt.Printf("%s %s %s %s %s",
		DbHost,
		DbName,
		DbUser,
		DbPassword,
		DbPort)

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%v",
		DbUser,
		DbPassword,
		DbHost,
		DbName)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&mod.User{})
	db.AutoMigrate(&cat.Category{})
	db.AutoMigrate(&cat.SubCategory{})
	db.AutoMigrate(&cat.Product{})

	db.AutoMigrate(&bis.Sale{})
	db.AutoMigrate((bis.SaleProduct{}))

	return db, err
}
