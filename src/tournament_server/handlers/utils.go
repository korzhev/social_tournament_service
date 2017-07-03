package handlers

import (
	"errors"
	"tournament_server/models"

	"fmt"

	"github.com/go-pg/pg"
)

const COUNT_SUM_SQL = "SELECT balance FROM money_transactions WHERE player_id = ? AND last_tx = TRUE FOR UPDATE"
const UPDATE_LAST_TX = "UPDATE money_transactions SET last_tx = FALSE WHERE player_id = ? AND last_tx = TRUE"

var LocalDB *pg.DB

type OkResponse struct {
	Message string `json:"message"`
}

type ResultMT struct {
	Balance uint64
}
type ResultUpdateLastMT struct {
	LastTX bool
}

func playerCanPay(tx *pg.Tx, pid string, points uint64) (uint64, error) {
	var result ResultMT
	_, err := tx.Query(&result, COUNT_SUM_SQL, pid)
	if err != nil {
		return result.Balance, err
	}
	if result.Balance == 0 {
		return result.Balance, errors.New("Player hasn't money, OR another transaction is being processed")
	}
	if result.Balance < points {
		return result.Balance, errors.New(fmt.Sprintf("Player has not enough money, only: %d", result.Balance))
	}
	return result.Balance, nil
}

func newMoneyTransaction(tx *pg.Tx, pid string, points uint64, transactionType uint8) (uint64, error) {
	var balance uint64
	var oldBalance uint64
	multi := multiplier(transactionType)
	if multi < 0 {
		current, err := playerCanPay(tx, pid, points)
		oldBalance = current
		if err != nil {
			return current, err
		}
		balance = current - points
	} else {
		var result ResultMT
		_, err := tx.Query(&result, COUNT_SUM_SQL, pid)
		if err != nil {
			return result.Balance, err
		}
		oldBalance = result.Balance
		balance = result.Balance + points
	}
	var result ResultUpdateLastMT

	_, err := tx.Query(&result, UPDATE_LAST_TX, pid)
	if err != nil {
		return oldBalance, err
	}
	mt := &models.MoneyTransaction{
		Type:     transactionType,
		Sum:      points,
		PlayerID: pid,
		Balance:  balance,
		LastTx:   true,
	}
	errMT := tx.Insert(mt)
	if errMT != nil {
		return oldBalance, errMT
	}
	return mt.Sum, nil
}

func multiplier(number uint8) int64 {
	if number%2 == 0 {
		return -1
	}
	return 1
}
