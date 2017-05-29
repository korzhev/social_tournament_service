package models

const ( // +(odd) add money, -(even) remove money
	FUND               uint8 = iota + 1 // +
	TAKE                                // -
	PRIZE                               // +
	BACKER_DONAT                        // -
	_                                   // +
	TOURNAMENT_DEPOSIT                  // -
	BACKER_PRIZE                        // +
)

type MoneyTransaction struct {
	ID       uint64
	PlayerID string
	Type     uint8
	Sum      int64
}
