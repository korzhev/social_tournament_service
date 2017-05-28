package main

import (
	"fmt"

	"tornament_server/config"

	"github.com/labstack/echo"
)

var CONF = config.GetConf()

func main() {
	e := echo.New()
	//DB := models.GetDB(CONF.DB)
	//defer DB.Close()
	//handlers.LocalDB = DB
	// routes
	//e.GET("/take", handlers.TakeHandler)
	//e.GET("/fund", handlers.FundHandler)
	//e.GET("/announceTournament", handlers.AnnounceHandler)
	//e.GET("/joinTournament", handlers.JoinHandler)
	//e.GET("/balance", handlers.BalanceHandler)
	//e.POST("/resultTournament", handlers.ResultHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", CONF.Port)))
}
