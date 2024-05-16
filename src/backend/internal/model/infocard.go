package model

import (
	"strconv"
	"time"
)

type InfoCardID int64

func ToInfoCardID(id int64) *InfoCardID {
	infoCardID := InfoCardID(id)
	return &infoCardID
}

func (i *InfoCardID) Int() int64 {
	return int64(*i)
}

func (i *InfoCardID) String() string {
	return strconv.FormatInt(i.Int(), 10)
}

type InfoCard struct {
	ID                *InfoCardID
	CreatedEmployeeID *EmployeeID
	IsConfirmed       bool
	CreatedDate       *time.Time
}

type FullInfoCard struct {
	ID                *InfoCardID
	CreatedEmployeeID *EmployeeID
	IsConfirmed       bool
	CreatedDate       *time.Time

	FullName    string
	PhoneNumber string
	CompanyID   *CompanyID
	Post        string
	DateOfBirth *time.Time
}
