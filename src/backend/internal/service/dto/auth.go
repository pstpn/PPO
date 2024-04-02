package dto

import (
	"course/internal/model"
	"time"
)

type RegisterEmployeeRequest struct {
	PhoneNumber string
	FullName    string
	CompanyID   int64
	Post        int64
	Password    *model.Password
	DateOfBirth *time.Time
}

type LoginEmployeeRequest struct {
	PhoneNumber string
	Password    string
}
