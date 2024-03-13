package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/logger"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) error
	LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) error
}

type authServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("encrypt password: %w", err)
	}

	request.Password = string(hashedPassword)
	err = a.employeeStorage.Register(ctx, request)
	if err != nil {
		return fmt.Errorf("create employee: %w", err)
	}

	return nil
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) error {
	user, err := a.employeeStorage.GetByPhone(ctx, &GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		return fmt.Errorf("get user by phone number: %w", err)
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
}
