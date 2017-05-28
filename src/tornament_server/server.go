package main

import (
	"fmt"

	"tornament_server/config"

	"tornament_server/handlers"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

var CONF = config.GetConf()

func main() {
	e := echo.New()
	db := pg.Connect(&pg.Options{
		User:     CONF.DB.User,
		Password: CONF.DB.Password,
		Database: CONF.DB.Database,
	})
	//DB := models.GetDB(CONF.DB)
	defer db.Close()
	handlers.LocalDB = db
	//routes
	e.GET("/take", handlers.TakeHandler)
	e.GET("/fund", handlers.FundHandler)
	e.GET("/announceTournament", handlers.AnnounceHandler)
	e.GET("/joinTournament", handlers.JoinHandler)
	e.GET("/balance", handlers.BalanceHandler)
	e.POST("/resultTournament", handlers.ResultHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", CONF.Port)))
}
