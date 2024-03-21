package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service/dto"
)

//go:generate mockery --name=CompanyStorage
type CompanyStorage interface {
	GetByID(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
}
