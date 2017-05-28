package handlers

import (
	"fmt"
	"net/http"

	"tornament_server/models"

	"github.com/labstack/echo"
)

func AnnounceHandler(c echo.Context) error {
	announce := new(Announce)
	if err := c.Bind(announce); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, AnnounceErrMsg}
	}
	at := &models.Tournament{
		TournamentID: announce.TournamentId,
		Deposit:      announce.Deposit,
	}
	err := LocalDB.Create(at).Error
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Tournament %d was announced with deposit: %d",
			announce.TournamentId,
			announce.Deposit)})
}

func JoinHandler(c echo.Context) error {
	join := new(Join)
	if err := c.Bind(join); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}

	// transaction
	tx := LocalDB.Begin()
	// find tournament
	tournament := &models.Tournament{}
	errT := tx.Where("tournament_id = ?", join.TournamentId).First(&tournament).Error
	if errT != nil {
		tx.Rollback()
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}

	// payment
	backersCount := uint64(len(join.Backers))
	if backersCount == 0 {
		// payment for 1 player
		_, err := newMoneyTransaction(tx, join.PlayerId, tournament.Deposit, models.TOURNAMENT_DEPOSIT)
		if err != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, err.Error()}
		}
	} else {
		// payment for player and backers
		paymentSum := tournament.Deposit / (backersCount + 1)
		// for player
		_, err := newMoneyTransaction(tx, join.PlayerId, paymentSum, models.TOURNAMENT_DEPOSIT)
		if err != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, err.Error()}
		}
		// for backers
		for _, player := range join.Backers {
			_, err := newMoneyTransaction(tx, player, paymentSum, models.BACKER_DONAT)
			if err != nil {
				tx.Rollback()
				return &echo.HTTPError{http.StatusBadRequest, err.Error()}
			}
		}
	}

	// join tournament
	je := &models.JoinEvent{
		TournamentID: join.TournamentId,
		PlayerId:     join.PlayerId,
		Backers:      join.Backers,
	}
	errJE := tx.Create(je).Error
	if errJE != nil {
		tx.Rollback()
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}
	tx.Commit()

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Player %d joined to tournament %d with backers: %v",
			join.PlayerId,
			join.TournamentId,
			join.Backers)})
}

func ResultHandler(c echo.Context) error {
	win := new(Win)
	if err := c.Bind(win); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, ResultErrMsg}
	}

	tx := LocalDB.Begin()
	for _, w := range win.Winners {
		je := &models.JoinEvent{}
		errJoins := tx.Where(
			"tournament_id = ? and player_id = ?",
			win.TournamentId,
			w.PlayerId).First(&je).Error
		if errJoins != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, errJoins.Error()}
		}
		prize := w.Prize / uint64(len(je.Backers)+1)
		_, errMT := newMoneyTransaction(tx, w.PlayerId, prize, models.PRIZE)

		if errMT != nil {
			tx.Rollback()
			return &echo.HTTPError{http.StatusBadRequest, errMT.Error()}
		}
		for _, b := range je.Backers {
			_, errBMT := newMoneyTransaction(tx, b, prize, models.BACKER_PRIZE)
			if errBMT != nil {
				tx.Rollback()
				return &echo.HTTPError{http.StatusBadRequest, errBMT.Error()}
			}
		}
	}

	errT := tx.Where("tournament_id = ?").Delete(models.Tournament{}).Error
	if errT != nil {
		tx.Rollback()
		return &echo.HTTPError{http.StatusBadRequest, errT.Error()}
	}

	tx.Commit()
	return c.JSON(
		http.StatusOK,
		&ResultResponse{win.Winners})
}
