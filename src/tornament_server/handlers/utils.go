package handlers

import (
	"errors"
	"strconv"
	"tornament_server/models"

	"fmt"

	"github.com/jinzhu/gorm"
)

const COUNT_SUM_SQL = "SELECT SUM(Sum) FROM money_transactions WHERE player_id = ?"

var LocalDB *gorm.DB

type OkResponse struct {
	Message string `json:"message"`
}

type ResultMT struct {
	Sum int64
}

func playerCanPay(tx *gorm.DB, pid string, points uint64) (int64, error) {
	var result ResultMT

	err := tx.Raw(COUNT_SUM_SQL, pid).Scan(&result).Error
	if err != nil {
		return result.Sum, errors.New(WrongPlayerMsg)
	}
	if result.Sum < int64(points) {
		return result.Sum, errors.New(fmt.Sprintf("Player has not enough money, only: %d", result.Sum))
	}
	return result.Sum, nil
}

func newMoneyTransaction(tx *gorm.DB, pid string, points uint64, transactionType uint8) (uint64, error) {
	sum, err := playerCanPay(tx, pid, points)
	if err != nil {
		return uint64(sum), err
	}
	mt := &models.MoneyTransaction{
		Type:     transactionType,
		Sum:      int64(points) * multiplier(transactionType),
		PlayerID: pid,
	}
	errMT := tx.Create(mt).Error
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

func getUint64Param(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

