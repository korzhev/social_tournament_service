package models

type Player struct {
	ID                uint64 `gorm:"primary_key"`
	MoneyTransactions []MoneyTransaction
}
