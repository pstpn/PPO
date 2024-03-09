package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	storage "course/internal/storage/postgres"
	"course/pkg/logger"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *RegisterEmployeeRequest) error
	LoginEmployee(ctx context.Context, request *LoginEmployeeRequest) error
}

type authServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage
}

type RegisterEmployeeRequest struct {
	PhoneNumber string
	FullName    string
	CompanyID   int64
	Post        int64
	Password    string
	DateOfBirth *time.Time
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *RegisterEmployeeRequest) error {
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

type LoginEmployeeRequest struct {
	PhoneNumber string
	Password    string
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *LoginEmployeeRequest) error {
	user, err := a.employeeStorage.GetByPhone(ctx, &GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		return fmt.Errorf("get user by phone number: %w", err)
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
}
