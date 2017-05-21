package main

import (
	"fmt"

	"tornament_server/config"
	"tornament_server/handlers"
	"tornament_server/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var CONF = config.GetConf()

var DB *gorm.DB

func main() {
	e := echo.New()
	DB = models.GetDB(CONF.DB)
	defer DB.Close()

	// routes
	e.GET("/take", handlers.TakeHandler)
	e.GET("/fund", handlers.FundHandler)
	e.GET("/announceTournament", handlers.AnnounceHandler)
	e.GET("/joinTournament", handlers.JoinHandler)
	e.GET("/balance", handlers.BalanceHandler)
	e.POST("/resultTournament", handlers.ResultHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", CONF.Port)))
}
