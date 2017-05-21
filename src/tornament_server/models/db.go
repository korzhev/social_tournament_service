package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB(params string) *gorm.DB {
	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Println("DB Connection Failure")
		log.Fatal(err)
	}
	db.AutoMigrate(&Player{}, &Tournament{}, &MoneyTransaction{})
	return db
}
