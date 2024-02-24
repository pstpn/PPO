package service

import (
	"context"

	"course/internal/model"
)

type FieldService interface {
	UpdateCardField(ctx *context.Context, request *UpdateCardFieldRequest) error
}

type UpdateCardFieldRequest struct {
	InfoCardID int64
	Field      *model.FieldType
	Value      string
}
