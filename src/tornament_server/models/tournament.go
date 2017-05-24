package models

type Tournament struct {
	ID uint64 `gorm:"primary_key"`
	TournamentAnnounces []TournamentAnnounce `gorm:"ForeignKey:TournamentID"`
}
