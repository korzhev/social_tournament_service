package models

type JoinEvent struct {
	ID           uint64 `gorm:"primary_key"`
	TournamentID string `gorm:"not null"`
	PlayerId     string `gorm:"not null"`
	Backers      []string
}
