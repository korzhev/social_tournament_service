package handlers

import (
	"fmt"
	"net/http"

	"tornament_server/models"

	"github.com/labstack/echo"
)

func TakeHandler(c echo.Context) error {
	pp := new(PlayerPoints)
	if err := c.Bind(pp); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}

	// transaction
	tx := LocalDB.Begin()
	result, err := newMoneyTransaction(tx, pp.PlayerId, pp.Points, models.TAKE)
	if err != nil {
		tx.Rollback()
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}
	tx.Commit()
	// end of transaction

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"%d points were taken from playerid: %d. %d left",
			pp.Points,
			pp.PlayerId,
			result)})
}

func FundHandler(c echo.Context) error {
	pp := new(PlayerPoints)
	if err := c.Bind(pp); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	_, err := newMoneyTransaction(LocalDB, pp.PlayerId, pp.Points, models.FUND)
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, WrongPlayerMsg}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Funded %d points to playerid: %d", pp.Points, pp.PlayerId)})
}

func BalanceHandler(c echo.Context) error {
	pid := c.QueryParam("playerId")
	if pid == "" {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}

	var result ResultMT

	err := LocalDB.Raw(COUNT_SUM_SQL, pid).Scan(&result).Error
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}

	return c.JSON(http.StatusOK, &BalanceResponse{PlayerId: pid, Balance: uint64(result.Sum)})
}
