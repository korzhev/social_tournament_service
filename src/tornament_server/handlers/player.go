package handlers

import (
	"fmt"
	"net/http"

	"tornament_server/models"

	"strconv"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const PlayerPointsErrMsg = "\"playerId\" and \"points\" are positive integers and required"

const PlayerBalanceErrMsg = "\"playerId\" is positive integer and required"

const COUNT_SUM_SQL = "SELECT SUM(Sum) FROM money_transactions WHERE player_id = ?"

type BalanceResponse struct {
	PlayerId string `json:"playerId"`
	Balance  uint64 `json:"balance"`
}

type ResultMT struct {
	Sum int64
}

func validationErrPidPoints(errPid error, errPoints error, pid uint64) bool {
	return errPid != nil || errPoints != nil || pid == 0
}

func takePointsFromPlayer(pid uint64, points uint64) (int64, error) {
	// transaction
	tx := LocalDB.Begin()

	var result ResultMT

	err := tx.Raw(COUNT_SUM_SQL, pid).Scan(&result).Error
	if err != nil {
		tx.Rollback()
		return result.Sum, errors.New(WrongPlayerMsg)
	}
	if result.Sum < int64(points) {
		tx.Rollback()
		return result.Sum, errors.New(fmt.Sprintf("Player has not enough money, only: %d", result.Sum))
	}
	mt := &models.MoneyTransaction{
		Type:     models.TAKE,
		Sum:      int64(points) * multiplier(models.TAKE),
		PlayerID: pid,
	}
	errMT := tx.Create(mt).Error
	if errMT != nil {
		return result.Sum, errors.New(WrongPlayerMsg)
	}

	tx.Commit()
	return result.Sum, nil
	// end of transaction
}

func TakeHandler(c echo.Context) error {
	pid, errPid := getUint64Param(c.QueryParam("playerId"))
	points, errPoints := getUint64Param(c.QueryParam("points"))
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}

	sum, err := takePointsFromPlayer(pid, points)
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"%d points were taken from playerid: %d. %d left",
			points,
			pid,
			uint64(sum)-points)})
}

func FundHandler(c echo.Context) error {
	pid, errPid := getUint64Param(c.QueryParam("playerId"))
	points, errPoints := getUint64Param(c.QueryParam("points"))
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}
	mt := &models.MoneyTransaction{
		Type:     models.FUND,
		Sum:      int64(points) * multiplier(models.FUND),
		PlayerID: pid,
	}
	err := LocalDB.Create(mt).Error
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, WrongPlayerMsg}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Funded %d points to playerid: %d", points, pid)})
}

func BalanceHandler(c echo.Context) error {
	pid, errPid := getUint64Param(c.QueryParam("playerId"))
	if validationErrPidPoints(errPid, nil, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}

	var result ResultMT

	err := LocalDB.Raw(COUNT_SUM_SQL, pid).Scan(&result).Error
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, PlayerBalanceErrMsg}
	}

	return c.JSON(http.StatusOK, &BalanceResponse{PlayerId: strconv.FormatUint(pid, 10), Balance: uint64(result.Sum)})
}
