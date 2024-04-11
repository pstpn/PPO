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

func NewAuthService(logger logger.Interface, employeeStorage storage.EmployeeStorage) AuthService {
	return &authServiceImpl{
		logger:          logger,
		employeeStorage: employeeStorage,
	}
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*model.Employee, error) {
	a.logger.Infof("register employee with phone %s", request.PhoneNumber)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password.Value), bcrypt.DefaultCost)
	if err != nil {
		a.logger.Errorf("encrypt password: %s", err.Error())
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
		a.logger.Errorf("create employee: %s", err.Error())
		return nil, fmt.Errorf("create employee: %w", err)
	}

	return employee, nil
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) error {
	a.logger.Infof("login employee with phone %s", request.PhoneNumber)

	user, err := a.employeeStorage.GetByPhone(ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		a.logger.Errorf("get user by phone number: %s", err.Error())
		return fmt.Errorf("get user by phone number: %w", err)
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password.Value), []byte(request.Password))
}
