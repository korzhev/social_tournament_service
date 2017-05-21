package models

import "time"

const ( // +(odd) add money, -(even) remove money
	FUND         uint8 = iota + 1 // +
	TAKE                          // -
	PRIZE                         // +
	BACKER_DONAT                  // -
	_                             // +
	PAYMENT                       // -
)

type MoneyTransaction struct {
	ID        uint64 `gorm:"primary_key"`
	Type      uint8
	Player    Player `gorm:"ForeignKey:PlayerID"`
	Sum       int64
	CreatedAt time.Time
}
