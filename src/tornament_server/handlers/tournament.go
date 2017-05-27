package handlers

import (
	"fmt"
	"net/http"

	"strings"

	"tornament_server/models"

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

func getBackers(str string) ([]uint64, error) {
	backersStr := strings.Split(str, ",")
	backers := []uint64{}

	for _, backer := range backersStr {
		id, err := getUint64Param(backer)
		if err != nil {
			return nil, err
		}
		backers = append(backers, id)
	}
	return backers, nil
}

func AnnounceHandler(c echo.Context) error {
	tournamentId, errTid := getUint64Param(c.QueryParam("tournamentId"))
	deposit, errDeposit := getUint64Param(c.QueryParam("deposit"))
	if validateId(errTid, tournamentId) || errDeposit != nil {
		return &echo.HTTPError{http.StatusBadRequest, AnnounceErrMsg}
	}

	at := &models.Tournament{
		TournamentID: tournamentId,
		Deposit:      deposit,
	}
	err := LocalDB.Create(at).Error
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf("Tournament %d was announced with deposit: %d", tournamentId, deposit)})
}

func JoinHandler(c echo.Context) error {
	pid, errPid := getUint64Param(c.QueryParam("playerId"))
	tournamentId, errTid := getUint64Param(c.QueryParam("tournamentId"))
	backers, err := getBackers(c.QueryParam("backerId"))
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}
	if validateId(errTid, tournamentId) || validateId(errPid, pid) {
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}

	tx := LocalDB.Begin()
	tournament := &models.Tournament{}
	tx.Where("tournament_id = ?", tournamentId).First(&tournament)

	backersCount := uint64(len(backers))

	if backersCount == 0 {
		// payment for 1 player
		_, err := takeMoney(tx, pid, tournament.Deposit, models.TOURNAMENT_DEPOSIT)
		if err != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, err.Error()}
		}
	} else {
		// payment for player and backers
		paymentSum := tournament.Deposit / (backersCount + 1)
		_, err := takeMoney(tx, pid, paymentSum, models.TOURNAMENT_DEPOSIT)
		if err != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, err.Error()}
		}
		for _, player := range backers {
			_, err := takeMoney(tx, player, paymentSum, models.BACKER_DONAT)
			if err != nil {
				tx.Rollback()
				return &echo.HTTPError{http.StatusBadRequest, err.Error()}
			}
		}
	}
	// join tournament
	je := &models.JoinEvent{
		TournamentID: tournamentId,
		PlayerId:     pid,
		Backers:      backers,
	}
	errJE := tx.Create(je).Error
	if errJE != nil {
		tx.Rollback()
		return &echo.HTTPError{http.StatusBadRequest, errJE.Error()}
	}

	tx.Commit()

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Player %d joined to tournament %d with backers: %v", pid, tournamentId, backers)})
}

func ResultHandler(c echo.Context) error {
	tournamentId, errTid := getUint64Param(c.FormValue("tournamentId"))
	if validateId(errTid, tournamentId) {
		return &echo.HTTPError{http.StatusBadRequest, ResultErrMsg}
	}
	winners := []PlayerPrize{{"1111", 100}, {"11211", 200}}
	return c.JSON(
		http.StatusOK,
		&ResultResponse{winners})
}
