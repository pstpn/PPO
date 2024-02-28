package service

import (
	"context"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *RegisterEmployeeRequest) error
	LoginEmployee(ctx context.Context, request *LoginEmployeeRequest) error
}

type RegisterEmployeeRequest struct {
	PhoneNumber string
	Password    string
}

type LoginEmployeeRequest struct {
	PhoneNumber string
	Password    string
}
