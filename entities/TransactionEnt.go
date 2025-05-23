package entities

import "time"

type Transaction struct {
	Id uint
	ConsumentId int
	ContractNumber string
	Otr int
	AdminFee int
	Installment int
	Interest int
	AssetName string
	CreatedAt time.Time
	UpdatedAt time.Time
}