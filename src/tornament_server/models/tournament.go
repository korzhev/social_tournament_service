package models

type Tournament struct {
	ID         uint64      `gorm:"primary_key"`
	Tournament uint64      `gorm:"not null"`
	Deposit    uint64      `gorm:"not null"`
	IsOpen     bool        `sql:"DEFAULT:TRUE"`
	JoinEvents []JoinEvent `gorm:"ForeignKey:TournamentID"`
}
