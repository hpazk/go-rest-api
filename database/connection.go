package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetInstance() *gorm.DB {
	dsn := "host=localhost user=postgres password= dbname=jersey_dev port=5432 sslmode=disable"

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	return db
}
