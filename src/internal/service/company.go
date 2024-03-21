package service

import (
	"context"
	"fmt"
	"log"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
)

type CompanyService interface {
	GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
}

type companyServiceImpl struct {
	logger         *log.Logger
	companyStorage storage.CompanyStorage
}

func (c *companyServiceImpl) GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error) {
	company, err := c.companyStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get company: %w", err)
	}

	return company, nil
}
