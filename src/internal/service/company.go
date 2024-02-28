package service

import (
	"context"
	"log"

	"course/internal/model"
)

type CompanyService interface {
	GetCompany(ctx context.Context, request *GetCompanyRequest) (*model.Company, error)
}

type GetCompanyRequest struct {
	CompanyID int64
}

// CompanyServiceImpl TODO
type CompanyServiceImpl struct {
	logger *log.Logger
	//storage CompanyRepository
}

//func (c *CompanyServiceImpl) GetCompany(ctx context.Context, request *GetCompanyRequest) (*model.Company, error) {
//	company, err := c.storage.GetCompanyByID(ctx, request)
//	if err != nil {
//		return nil, fmt.Errorf("get company by ID: %w", err)
//	}
//
//	return company, nil
//}
