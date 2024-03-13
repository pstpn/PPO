package dto

import (
	"time"
)

type RegisterEmployeeRequest struct {
	PhoneNumber string
	FullName    string
	CompanyID   int64
	Post        int64
	Password    string
	DateOfBirth *time.Time
}

type LoginEmployeeRequest struct {
	PhoneNumber string
	Password    string
}
