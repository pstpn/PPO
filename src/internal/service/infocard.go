package service

import (
	"context"
	"time"

	"course/internal/model"
)

type InfoCardService interface {
	CreateInfoCard(ctx *context.Context, request *CreateInfoCardRequest) error
	ListAllInfoCards(ctx *context.Context, request *ListAllInfoCardsRequest) []*model.InfoCard
	DeleteInfoCard(ctx *context.Context, request *DeleteCardFieldRequest) error
}

type CreateInfoCardRequest struct {
	PhoneNumber   string
	FullName      string
	Birthday      *time.Time
	EmploymentDay *time.Time
}

// ListAllInfoCardsRequest TODO: pagination, sort, filter
type ListAllInfoCardsRequest struct {
}

type DeleteCardFieldRequest struct {
	InfoCardID int64
}
