package models

type JoinEvent struct {
	ID           uint64
	TournamentID string
	PlayerId     string
	Backers      []string
}
