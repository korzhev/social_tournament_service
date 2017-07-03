package handlers

import (
	"fmt"
	"net/http"

	"tournament_server/models"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

func TakeHandler(c echo.Context) error {
	pp := new(PlayerPoints)
	if err := c.Bind(pp); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	var money uint64
	// transaction
	err := LocalDB.RunInTransaction(func(tx *pg.Tx) error {
		result, err := newMoneyTransaction(tx, pp.PlayerId, pp.Points, models.TAKE)
		if err != nil {
			return err
		}
		money = result
		return nil
	})
	// end of transaction

	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"%d points were taken from playerid: %v.",
			pp.Points,
			pp.PlayerId)})
}

func FundHandler(c echo.Context) error {
	pp := new(PlayerPoints)
	if err := c.Bind(pp); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}

	var money uint64
	// transaction
	err := LocalDB.RunInTransaction(func(tx *pg.Tx) error {
		result, err := newMoneyTransaction(tx, pp.PlayerId, pp.Points, models.FUND)
		if err != nil {
			return err
		}
		money = result
		return nil
	})
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Funded %v points to playerid: %v", pp.Points, pp.PlayerId)})
}

func BalanceHandler(c echo.Context) error {
	pid := c.QueryParam("playerId")
	if pid == "" {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}

	var result ResultMT
	_, err := LocalDB.Query(&result, COUNT_SUM_SQL, pid)
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(http.StatusOK, &BalanceResponse{PlayerId: pid, Balance: result.Balance})
}
