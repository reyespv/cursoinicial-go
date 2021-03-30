package main

import (
	"log"
	"os"

	conf "./conf"
	"github.com/joho/godotenv"
)

func main() {
	//obtenemos los parametros de conexion
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
		panic("failed to resolve env file... ")
	}

	//obtenemos una interfaz o referencia hacia la conexion DB
	_, errDb := conf.ConnectDB(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))

	if errDb != nil {
		panic(errDb)
	}

}
