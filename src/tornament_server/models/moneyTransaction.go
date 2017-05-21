package models

const ( // +(odd) add money, -(even) remove money
	FUND         uint8 = iota + 1 // +
	TAKE                          // -
	PRIZE                         // +
	BACKER_DONAT                  // -
	_                             // +
	PAYMENT                       // -
)

type MoneyTransaction struct {
	ID       uint64 `gorm:"primary_key"`
	PlayerID uint64 `gorm:"index;not null"`
	Type     uint8  `gorm:"not null"`
	Sum      int64	`gorm:"not null"`
}
