package handlers

import (
	"fmt"
	"net/http"

	"tournament_server/models"

	"github.com/go-pg/pg"
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
	err := LocalDB.Insert(at)
	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Tournament %v was announced with deposit: %v",
			announce.TournamentId,
			announce.Deposit)})
}

func JoinHandler(c echo.Context) error {
	join := new(Join)
	if err := c.Bind(join); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, JoinErrMsg}
	}

	tournament := &models.Tournament{}
	backersCount := uint64(len(join.Backers))
	// transaction
	err := LocalDB.RunInTransaction(func(tx *pg.Tx) error {
		// find tournament
		errT := tx.Model(&tournament).Column("*").Where("tournament_id = ?", join.TournamentId).Select()
		if errT != nil {
			return errT
		}
		paymentSum := tournament.Deposit / (backersCount + 1)

		// take mone from player
		_, err := newMoneyTransaction(tx, join.PlayerId, paymentSum, models.TOURNAMENT_DEPOSIT)
		if err != nil {
			return err
		}

		// toke money from backers
		if backersCount != 0 {
			for _, player := range join.Backers {
				_, err := newMoneyTransaction(tx, player, paymentSum, models.BACKER_DONAT)
				if err != nil {
					return err
				}
			}
		}

		// save player join event
		je := &models.JoinEvent{
			TournamentID: join.TournamentId,
			PlayerId:     join.PlayerId,
			Backers:      join.Backers,
		}

		errJE := tx.Insert(je)
		if errJE != nil {
			return errJE
		}

		return nil
	})
	// end of transaction

	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&OkResponse{Message: fmt.Sprintf(
			"Player %v joined to tournament %v with backers: %v",
			join.PlayerId,
			join.TournamentId,
			join.Backers)})
}

func ResultHandler(c echo.Context) error {
	win := new(Win)
	if err := c.Bind(win); err != nil {
		return &echo.HTTPError{http.StatusBadRequest, ResultErrMsg}
	}

	err := LocalDB.RunInTransaction(func(tx *pg.Tx) error {
		// for each winner
		for _, w := range win.Winners {
			// find winner join event
			je := &models.JoinEvent{}
			errJoins := tx.Model(&je).Column("*").Where(
				"tournament_id = ? and player_id = ?",
				win.TournamentId,
				w.PlayerId).Select()
			if errJoins != nil {
				return errJoins
			}

			// count prize
			prize := w.Prize / uint64(len(je.Backers)+1)

			// give prize to winner
			_, errMT := newMoneyTransaction(tx, w.PlayerId, prize, models.PRIZE)
			if errMT != nil {
				return errMT
			}

			// give bonus for each backer
			for _, b := range je.Backers {
				_, errBMT := newMoneyTransaction(tx, b, prize, models.BACKER_PRIZE)
				if errBMT != nil {
					return errBMT
				}
			}
		}

		// delete announced tournament
		_, errT := tx.Model(&models.Tournament{}).Where("tournament_id = ?", win.TournamentId).Delete()
		if errT != nil {
			return errT
		}

		return nil
	})

	if err != nil {
		return &echo.HTTPError{http.StatusBadRequest, err.Error()}
	}

	return c.JSON(
		http.StatusOK,
		&ResultResponse{win.Winners})
}
