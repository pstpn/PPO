package service

import (
	"context"
	"fmt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type CompanyService interface {
	Create(ctx context.Context, request *dto.CreateCompanyRequest) (*model.Company, error)
	GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error)
}

type companyServiceImpl struct {
	logger         logger.Interface
	companyStorage storage.CompanyStorage
}

func NewCompanyService(logger logger.Interface, companyStorage storage.CompanyStorage) CompanyService {
	return &companyServiceImpl{
		logger:         logger,
		companyStorage: companyStorage,
	}
}

func (c *companyServiceImpl) Create(ctx context.Context, request *dto.CreateCompanyRequest) (*model.Company, error) {
	company, err := c.companyStorage.Create(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("create company: %w", err)
	}

	return company, nil
}

func (c *companyServiceImpl) GetCompany(ctx context.Context, request *dto.GetCompanyRequest) (*model.Company, error) {
	company, err := c.companyStorage.GetByID(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("get company: %w", err)
	}

	return company, nil
}
