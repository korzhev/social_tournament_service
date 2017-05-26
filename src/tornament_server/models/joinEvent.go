package models

type JoinEvent struct {
	ID           uint64   `gorm:"primary_key"`
	TournamentID string   `gorm:"index;not null;unique"`
	PlayerId     uint64   `gorm:"not null"`
	Backers      []Backer `gorm:"ForeignKey:PlayerID"`
}
