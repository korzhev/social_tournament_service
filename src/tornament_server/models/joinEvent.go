package models

type JoinEvent struct {
	ID           uint64 `gorm:"primary_key"`
	TournamentID uint64 `gorm:"index;not null;unique"`
	PlayerId     uint64 `gorm:"not null;unique"`
	Backers      []uint64
}
