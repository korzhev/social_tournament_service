package handlers

import (
	"strconv"
	"tornament_server/models"
)

type OkResponse struct {
	Message string `json:"message"`
}

func multiplier(number uint8) int64{
	if number%2 == 0 {
		return 1
	}
	return -1
}

func getUintParam(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func validateId(errPid error, pid uint64) bool {
	return errPid != nil || pid == 0
}

func countMoney(mt []models.MoneyTransaction) uint64 {
	var money int64
	for _, t := range mt {
		money += t.Sum
	}
	return uint64(money)
}
