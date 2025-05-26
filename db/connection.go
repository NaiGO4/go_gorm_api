package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal("Error al conectar a la base de datos: ", error)
	} else {
		log.Println("Conexion realizada exitosamente a la base de datos")
	}
}
