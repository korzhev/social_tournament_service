package handlers

type Announce struct {
	TournamentId string `json:"tournamentId" form:"tournamentId" query:"tournamentId"`
	Deposit      uint64 `json:"deposit" form:"deposit" query:"deposit"`
}

type Win struct {
	TournamentId string        `json:"tournamentId" form:"tournamentId" query:"tournamentId"`
	Winners      []PlayerPrize `json:"winners" form:"winners" query:"winners"`
}

type PlayerPrize struct {
	PlayerId string `json:"playerId"`
	Prize    uint64 `json:"prize"`
}


type ResultResponse struct {
	Winners []PlayerPrize `json:"winners"`
}

type Join struct {
	TournamentId string   `json:"tournamentId" form:"tournamentId" query:"tournamentId"`
	PlayerId     string   `json:"playerId" form:"playerId" query:"playerId"`
	Backers      []string `json:"backerId" form:"backerId" query:"backerId"`
}
