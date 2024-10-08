package database

import (
	"log"

	"github.com/Angxandralol/Proxy-App/server/utilities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DSN := utilities.Lector("CONNSTR")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Se Conecto a la DB")
	}
}
