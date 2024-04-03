package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*model.Employee, error)
	LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) error
}

type authServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*model.Employee, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password.Value), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("encrypt password: %w", err)
	}

	employee, err := a.employeeStorage.Register(ctx, &dto.RegisterEmployeeRequest{
		PhoneNumber: request.PhoneNumber,
		FullName:    request.FullName,
		CompanyID:   request.CompanyID,
		Post:        request.Post,
		Password: &model.Password{
			Value:    string(hashedPassword),
			IsHashed: true,
		},
		DateOfBirth: request.DateOfBirth,
	})
	if err != nil {
		return nil, fmt.Errorf("create employee: %w", err)
	}

	return employee, nil
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) error {
	user, err := a.employeeStorage.GetByPhone(ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		return fmt.Errorf("get user by phone number: %w", err)
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password.Value), []byte(request.Password))
}
