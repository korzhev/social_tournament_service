package models

type Tournament struct {
	ID           uint
	TournamentID string
	Deposit      uint64
	JoinEvents   []*JoinEvent
}
