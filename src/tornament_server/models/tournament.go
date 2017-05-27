package models

type Tournament struct {
	ID           uint        `gorm:"primary_key"`
	TournamentID uint64      `gorm:"not null;unique"`
	Deposit      uint64      `gorm:"not null"`
	JoinEvents   []JoinEvent `gorm:"ForeignKey:TournamentID"`
}
