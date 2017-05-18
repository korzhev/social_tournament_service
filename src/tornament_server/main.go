package main

import (
	"net/http"

	"strconv"

	"fmt"

	"github.com/labstack/echo"
)

type OkResponse struct {
	Message string `json:"message" form:"message" query:"message"`
}

const PlayerPointsErrMsg = "\"playerId\" and \"points\" are positive integers and required"

func validationErrPidPoints(errPid error, errPoints error, pid uint64) bool {
	return errPid != nil || errPoints != nil || pid == 0
}

func takeHandler(c echo.Context) error {
	pidStr := c.QueryParam("playerId")
	pid, errPid := strconv.ParseUint(pidStr, 10, 64)
	pointsStr := c.QueryParam("points")
	points, errPoints := strconv.ParseUint(pointsStr, 10, 64)
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	return c.JSON(http.StatusOK, &OkResponse{Message: fmt.Sprintf("%d points were taken from playerid: %d", points, pid)})
}
func fundHandler(c echo.Context) error {
	pidStr := c.QueryParam("playerId")
	pid, errPid := strconv.ParseUint(pidStr, 10, 64)
	pointsStr := c.QueryParam("points")
	points, errPoints := strconv.ParseUint(pointsStr, 10, 64)
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	return c.JSON(http.StatusOK, &OkResponse{Message: fmt.Sprintf("Funded %d points to playerid: %d", points, pid)})
}

func main() {
	e := echo.New()
	// routes
	e.GET("/take", takeHandler)
	e.GET("/fund", fundHandler)
	e.GET("/announceTournament", fundHandler)
	e.GET("/joinTournament", fundHandler)
	e.GET("/balance", fundHandler)
	e.POST("/resultTournament", fundHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
