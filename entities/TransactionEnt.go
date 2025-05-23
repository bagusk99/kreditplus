package entities

import (
	"time"
)

type Transaction struct {
	Id uint
	ConsumentId int
	Consument Consument
	ContractNumber string
	Otr int
	AdminFee int
	Installment int
	Interest int
	AssetName string
	CreatedAt time.Time
	UpdatedAt time.Time
}