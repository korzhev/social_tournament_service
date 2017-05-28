package handlers

type BalanceResponse struct {
	PlayerId string `json:"playerId"`
	Balance  uint64 `json:"balance"`
}

type PlayerPoints struct {
	PlayerId string `json:"playerId" form:"playerId" query:"playerId"`
	Points   uint64 `json:"points" form:"points" query:"points"`
}
