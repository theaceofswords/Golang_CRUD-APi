package config

import (
	"fmt"
	"golang-training/app/models"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	//host = "docker-postgres" // while using docker
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "user123"
	dbname   = "datb1"
)

var (
	db  *gorm.DB
	err error
)

func OpenDB() *gorm.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Employee{})

	return db
}
