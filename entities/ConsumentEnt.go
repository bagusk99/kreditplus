package entities

import "time"

type Consument struct {
	Id uint
	Nik string
	FullName string
	LegalName string
	PlaceOfBirth string
	DateOfBirth time.Time
	Salary int
	KtpPhoto string
	SelfiePhoto string
	CreatedAt time.Time
	UpdatedAt time.Time
}