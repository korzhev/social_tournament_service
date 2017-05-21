package handlers

import (
	"fmt"
	"net/http"

	"tornament_server/models"

	"github.com/labstack/echo"
)

const PlayerPointsErrMsg = "\"playerId\" and \"points\" are positive integers and required"

const PlayerBalanceErrMsg = "\"playerId\" is positive integer and required"

type BalanceResponse struct {
	PlayerId string `json:"playerId"`
	Balance  uint64 `json:"balance"`
}

type ResultMT struct {
	Score int64
}

func validationErrPidPoints(errPid error, errPoints error, pid uint64) bool {
	return errPid != nil || errPoints != nil || pid == 0
}

func TakeHandler(c echo.Context) error {
	pid, errPid := getUint64Param(c.QueryParam("playerId"))
	points, errPoints := getUint64Param(c.QueryParam("points"))
	if validationErrPidPoints(errPid, errPoints, pid) {
		return &echo.HTTPError{http.StatusBadRequest, PlayerPointsErrMsg}
	}

	//playerMT := []models.MoneyTransaction{}
	//mt := &models.MoneyTransaction{
	//	Type:     models.TAKE,
	//	Sum:      int64(points) * multiplier(models.TAKE),
	//	PlayerID: pid,
	//}
	// transaction
	tx := LocalDB.Begin()
	//var score []int64
	type Result struct {
		Score int64
	}
	var result []uint8
	row := tx.Model(&models.MoneyTransaction{}).Where("player_id = ?", pid).Select(
		"sum(sum)").Row()
	//defer rows.Close()
	err := row.Scan(result).Error()
	fmt.Println(result, err)
	//if err != nil {
	//	fmt.Println(err.Error)
	//	tx.Rollback()
	//	return &echo.HTTPError{http.StatusBadRequest, "wrong"}
	//}
	//fmt.Println(row)

	tx.Commit()

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("%d points were taken from playerid: %d", points, pid)})
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
	var balance uint64 = 100

	return c.JSON(http.StatusOK, &BalanceResponse{PlayerId: string(pid), Balance: balance})
}
