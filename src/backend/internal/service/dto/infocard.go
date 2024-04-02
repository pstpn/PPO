package dto

import (
	"time"
)

type CreateInfoCardRequest struct {
	EmployeeID  int64
	IsConfirmed bool
	CreatedDate *time.Time
}

type ValidateInfoCardRequest struct {
	InfoCardID  int64
	IsConfirmed bool
}

type GetInfoCardRequest struct {
	InfoCardID int64
}

// ListInfoCardsRequest TODO: pagination, sort, filter
type ListInfoCardsRequest struct {
}

type DeleteInfoCardRequest struct {
	InfoCardID int64
}
