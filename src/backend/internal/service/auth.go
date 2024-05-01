package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/jwt"
	"course/pkg/logger"
)

type AuthService interface {
	RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*tokens, error)
	LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) (*tokens, error)
	VerifyEmployeeAccessToken(ctx context.Context, request *dto.VerifyEmployeeAccessTokenRequest) (string, error)
	RefreshTokens(ctx context.Context, request *dto.RefreshEmployeeTokensRequest) (*tokens, error)
}

type authServiceImpl struct {
	logger          logger.Interface
	employeeStorage storage.EmployeeStorage

	tokenManager    jwt.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(
	logger logger.Interface,
	employeeStorage storage.EmployeeStorage,
	tokenManager jwt.TokenManager,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) AuthService {
	return &authServiceImpl{
		logger:          logger,
		employeeStorage: employeeStorage,

		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

type tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *authServiceImpl) RegisterEmployee(ctx context.Context, request *dto.RegisterEmployeeRequest) (*tokens, error) {
	a.logger.Infof("register employee with phone %s", request.PhoneNumber)

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	refreshTokenExpiredAt := time.Now().Add(a.refreshTokenTTL)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		a.logger.Errorf("hash password: %s", err.Error())
		return nil, fmt.Errorf("hash password: %w", err)
	}

	employee, err := a.employeeStorage.Register(ctx, &dto.RegisterEmployeeRequest{
		PhoneNumber:    request.PhoneNumber,
		FullName:       request.FullName,
		CompanyID:      request.CompanyID,
		Post:           request.Post,
		Password:       string(hashedPassword),
		RefreshToken:   refreshToken,
		TokenExpiredAt: &refreshTokenExpiredAt,
		DateOfBirth:    request.DateOfBirth,
	})
	if err != nil {
		a.logger.Errorf("create employee: %s", err.Error())
		return nil, fmt.Errorf("create employee: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(employee.PhoneNumber, time.Now())
	if err != nil {
		a.logger.Errorf("create access token for employee: %s", err.Error())
		return nil, fmt.Errorf("create access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authServiceImpl) LoginEmployee(ctx context.Context, request *dto.LoginEmployeeRequest) (*tokens, error) {
	a.logger.Infof("login employee with phone %s", request.PhoneNumber)

	employee, err := a.employeeStorage.GetByPhone(ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		a.logger.Errorf("get employee: %s", err.Error())
		return nil, fmt.Errorf("get employee: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password: %w", err)
	}

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	refreshTokenExpiredAt := time.Now().Add(a.refreshTokenTTL)

	err = a.employeeStorage.UpdateRefreshToken(ctx, &dto.UpdateToken{
		PhoneNumber:    employee.PhoneNumber,
		RefreshToken:   refreshToken,
		TokenExpiredAt: &refreshTokenExpiredAt,
	})
	if err != nil {
		a.logger.Errorf("update employee refresh token: %s", err.Error())
		return nil, fmt.Errorf("update employee refresh token: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(employee.PhoneNumber, time.Now())
	if err != nil {
		a.logger.Errorf("update access token for employee: %s", err.Error())
		return nil, fmt.Errorf("update access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authServiceImpl) VerifyEmployeeAccessToken(ctx context.Context, request *dto.VerifyEmployeeAccessTokenRequest) (string, error) {
	a.logger.Infof("verify employee access token")

	phoneNumber, err := a.tokenManager.Parse(request.AccessToken, a.accessTokenTTL)
	if err != nil {
		// Unwrap, because we need to check `github.com/golang-jwt/jwt/v5` library errors
		// using `errors.Is()` method
		return phoneNumber, err
	}

	return phoneNumber, nil
}

func (a *authServiceImpl) RefreshTokens(ctx context.Context, request *dto.RefreshEmployeeTokensRequest) (*tokens, error) {
	a.logger.Infof("refresh employee tokens")

	employee, err := a.employeeStorage.GetByPhone(ctx, &dto.GetEmployeeRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		a.logger.Errorf("get employee: %s", err.Error())
		return nil, fmt.Errorf("get employee: %w", err)
	}
	if employee.PhoneNumber != request.PhoneNumber {
		a.logger.Errorf("invalid phone number")
		return nil, fmt.Errorf("invalid phone number")
	}
	if employee.RefreshToken != request.RefreshToken {
		a.logger.Errorf("invalid refresh token")
		return nil, fmt.Errorf("invalid refresh token")
	}

	refreshToken, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		a.logger.Errorf("create refresh token for employee: %s", err.Error())
		return nil, fmt.Errorf("create refresh token for employee: %w", err)
	}
	tokenExpiredAt := time.Now().Add(a.refreshTokenTTL)

	err = a.employeeStorage.UpdateRefreshToken(ctx, &dto.UpdateToken{
		PhoneNumber:    employee.PhoneNumber,
		RefreshToken:   refreshToken,
		TokenExpiredAt: &tokenExpiredAt,
	})
	if err != nil {
		a.logger.Errorf("update employee refresh token: %s", err.Error())
		return nil, fmt.Errorf("update employee refresh token: %w", err)
	}

	accessToken, err := a.tokenManager.NewJWT(employee.PhoneNumber, time.Now())
	if err != nil {
		a.logger.Errorf("update access token for employee: %s", err.Error())
		return nil, fmt.Errorf("update access token for employee: %w", err)
	}

	return &tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
