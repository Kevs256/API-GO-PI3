package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection(IP_DB, PORT_DB, USER_DB, PASS_DB, NAME_DB string) {

	var DBstring = "host=" + IP_DB + " port=" + PORT_DB + " user=" + USER_DB + " password=" + PASS_DB + " dbname=" + NAME_DB

	var error error
	DB, error = gorm.Open(postgres.Open(DBstring), &gorm.Config{})
	if error != nil {
		fmt.Print("no se conectó")
	} else {
		log.Println("conexion exitosa")
	}
}
