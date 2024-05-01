package dto

import (
	"time"

	"course/pkg/storage/postgres"
)

type CreateInfoCardRequest struct {
	EmployeePhoneNumber string
	IsConfirmed         bool
	CreatedDate         *time.Time
}

type ValidateInfoCardRequest struct {
	InfoCardID  int64
	IsConfirmed bool
}

type GetInfoCardRequest struct {
	InfoCardID int64
}

type ListInfoCardsRequest struct {
	Pagination *postgres.Pagination
}

type DeleteInfoCardRequest struct {
	InfoCardID int64
}
