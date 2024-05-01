package model

import "time"

type InfoCardID int64

func ToInfoCardID(id int64) *InfoCardID {
	infoCardID := InfoCardID(id)
	return &infoCardID
}

func (i *InfoCardID) Int() int64 {
	return int64(*i)
}

type InfoCard struct {
	ID                         *InfoCardID
	CreatedEmployeePhoneNumber string
	IsConfirmed                bool
	CreatedDate                *time.Time
}
