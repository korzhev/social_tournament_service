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
	//db.LogMode(false)
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	//var mt MoneyTransaction
	var tournament Tournament
	var je JoinEvent

	db.Model(&tournament).Related(&je)

	// Auto creating FK not working!
	//db.AutoMigrate(&tournament, &mt, &je)

	return db
}
