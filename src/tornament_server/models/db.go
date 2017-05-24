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
	var player Player
	var mt MoneyTransaction
	var tournament Tournament
	var ta TournamentAnnounce
	db.Model(&player).Related(&mt)
	db.Model(&tournament).Related(&ta)
	// Auto creating FK not working!
	db.AutoMigrate(&player, &tournament, &mt, &ta)
	return db
}
