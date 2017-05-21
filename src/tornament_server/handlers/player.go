package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const PlayerPointsErrMsg = "\"playerId\" and \"points\" are positive integers and required"

const PlayerBalanceErrMsg = "\"playerId\" is positive integer and required"

type BalanceResponse struct {
	PlayerId string `json:"playerId"`
	Balance  uint64 `json:"balance"`
}

func validationErrPidPoints(errPid error, errPoints error, pid uint64) bool {
	return errPid != nil || errPoints != nil || pid == 0
}

func TakeHandler(c echo.Context) error {
	pid, errPid := getUintParam(c.QueryParam("playerId"))
	points, errPoints := getUintParam(c.QueryParam("points"))
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("%d points were taken from playerid: %d", points, pid)})
}

func FundHandler(c echo.Context) error {
	pid, errPid := getUintParam(c.QueryParam("playerId"))
	points, errPoints := getUintParam(c.QueryParam("points"))
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Funded %d points to playerid: %d", points, pid)})
}

func BalanceHandler(c echo.Context) error {
	pid, errPid := getUintParam(c.QueryParam("playerId"))
	if validationErrPidPoints(errPid, nil, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}
	var balance uint64 = 100
	return c.JSON(http.StatusOK, &BalanceResponse{PlayerId: string(pid), Balance: balance})
}
