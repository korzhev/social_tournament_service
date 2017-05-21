package handlers

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/labstack/echo"
)

const AnnounceErrMsg = "\"tournamentId\" and \"deposit\" are positive integers and required"
const JoinErrMsg = "\"tournamentId\", \"backerId\" and \"playerId\" are positive integers, " +
	"\"tournamentId\" and \"playerId\" are required"
const ResultErrMsg = "\"tournamentId\" is positive integer and required"

type PlayerPrize struct {
	PlayerId string `json:"playerId"`
	Prise    uint64 `json:"balance"`
}

type ResultResponse struct {
	Winners []PlayerPrize `json:"winners"`
}

func AnnounceHandler(c echo.Context) error {
	tournamentId, errTid := getUintParam(c.QueryParam("tournamentId"))
	deposit, errDeposit := getUintParam(c.QueryParam("deposit"))
	if validateId(errTid, tournamentId) || errDeposit != nil {
		return &echo.HTTPError{http.StatusBadRequest, AnnounceErrMsg}
	}
	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Tournament %d was announced with deposit: %d", tournamentId, deposit)})
}

func JoinHandler(c echo.Context) error {
	pid, errPid := getUintParam(c.QueryParam("playerId"))
	tournamentId, errTid := getUintParam(c.QueryParam("tournamentId"))
	backersStr := strings.Split(c.QueryParam("backerId"), ",")
	backers := []uint64{}
	for _, backer := range backersStr {
		id, err := getUintParam(backer)
		if err != nil {
			return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
		}
		backers = append(backers, id)
	}
	if validateId(errTid, tournamentId) || validateId(errPid, pid) {
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}
	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Player %d joined to tournament %d with backers: %v", pid, tournamentId, backers)})
}

func ResultHandler(c echo.Context) error {
	tournamentId, errTid := getUintParam(c.FormValue("tournamentId"))
	if validateId(errTid, tournamentId) {
		return &echo.HTTPError{http.StatusBadRequest, ResultErrMsg}
	}
	winners := []PlayerPrize{{"1111", 100}, {"11211", 200}}
	return c.JSON(
		http.StatusOK,
		&ResultResponse{winners})
}
