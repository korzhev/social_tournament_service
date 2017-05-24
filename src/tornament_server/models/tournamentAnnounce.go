package models

type TournamentAnnounce struct {
	ID           uint64 `gorm:"primary_key"`
	TournamentID uint64 `gorm:"index;not null"`
	Deposit      uint64 `gorm:"not null"`
}
