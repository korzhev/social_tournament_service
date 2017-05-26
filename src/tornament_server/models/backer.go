package models

type Backer struct {
	ID       uint64 `gorm:"primary_key"`
	PlayerID uint64 `gorm:"index;not null"`
}
