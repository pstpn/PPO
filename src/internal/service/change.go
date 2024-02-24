package service

import (
	"context"

	"course/internal/model"
)

type ChangeService interface {
	CreateChange(ctx *context.Context, request *CreateChangeRequest) error
	DeleteChange(ctx *context.Context, request *DeleteChangeRequest) error
}

type CreateChangeRequest struct {
	EmployeeID *model.EmployeeID
	FieldID    *model.InfoCardID
	OldValue   string
	NewValue   string
}

type DeleteChangeRequest struct {
	ID *model.ChangeID
}
