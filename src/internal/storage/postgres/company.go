package storage

import (
	"context"

	"course/internal/model"
	"course/internal/service"
	"course/pkg/storage/postgres"
)

type CompanyStorage interface {
	GetByID(ctx context.Context, request *service.GetCompanyRequest) (*model.Company, error)
}

type companyStorageImpl struct {
	db *postgres.Postgres
}
