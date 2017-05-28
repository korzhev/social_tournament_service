package handlers

import (
	"errors"
	"tornament_server/models"

	"fmt"

	"github.com/go-pg/pg"
)

const COUNT_SUM_SQL = "SELECT SUM(Sum) FROM money_transactions WHERE player_id = ?"

var LocalDB *pg.DB

type OkResponse struct {
	Message string `json:"message"`
}

type ResultMT struct {
	Sum int64
}

func playerCanPay(tx *pg.Tx, pid string, points uint64) (int64, error) {
	var result ResultMT
	_, err := tx.QueryOne(&result, COUNT_SUM_SQL, pid)
	if err != nil {
		return result.Sum, err
	}
	if result.Sum < int64(points) {
		return result.Sum, errors.New(fmt.Sprintf("Player has not enough money, only: %d", result.Sum))
	}
	return result.Sum, nil
}

func newMoneyTransaction(tx *pg.Tx, pid string, points uint64, transactionType uint8) (uint64, error) {
	var sum int64
	multi := multiplier(transactionType)
	if multi < 0 {
		current, err := playerCanPay(tx, pid, points)
		sum = current
		if err != nil {
			return uint64(sum), err
		}
	}

	mt := &models.MoneyTransaction{
		Type:     transactionType,
		Sum:      int64(points) * multi,
		PlayerID: pid,
	}
	errMT := tx.Insert(mt)
	if errMT != nil {
		return uint64(sum), errMT
	}
	return uint64(sum + mt.Sum), nil
}

func multiplier(number uint8) int64 {
	if number%2 == 0 {
		return -1
	}
	return 1
}
