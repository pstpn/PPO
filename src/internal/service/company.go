package service

import (
	"context"
	"fmt"
	"log"

	"course/internal/model"
	storage "course/internal/storage/postgres"
)

type CompanyService interface {
	GetCompany(ctx context.Context, request *GetCompanyRequest) (*model.Company, error)
}

type companyServiceImpl struct {
	logger         *log.Logger
	companyStorage storage.CompanyStorage
}

type GetCompanyRequest struct {
	CompanyID int64
}

func (c *companyServiceImpl) GetCompany(ctx context.Context, request *GetCompanyRequest) (*model.Company, error) {
	company, err := c.companyStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get company: %w", err)
	}

	return company, nil
}
