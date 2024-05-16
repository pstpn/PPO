package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	httputils "course/internal/controller/http/utils"
	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

type AuthController struct {
	l           logger.Interface
	authService service.AuthService
}

func NewAuthController(l logger.Interface, authService service.AuthService) *AuthController {
	return &AuthController{
		l:           l,
		authService: authService,
	}
}

type registerRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	CompanyID   int64  `json:"companyID"`
	Post        string `json:"post"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (a *AuthController) Register(c *gin.Context) {
	httputils.DisableCors(c)

	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	birthDate, err := time.Parse("02.01.2006", req.DateOfBirth)
	if err != nil {
		a.l.Errorf("incorrect request body date of birth: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body date of birth"})
		return
	}

	tokens, err := a.authService.RegisterEmployee(c.Request.Context(), &dto.RegisterEmployeeRequest{
		PhoneNumber:    req.PhoneNumber,
		FullName:       req.Name + " " + req.Surname,
		CompanyID:      req.CompanyID,
		Post:           model.ToPostTypeFromString(req.Post).Int(),
		Password:       req.Password,
		RefreshToken:   "",
		TokenExpiredAt: nil,
		DateOfBirth:    &birthDate,
	})
	if err != nil {
		err = fmt.Errorf("can`t register employee: %w", err)
		a.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
		"isAdmin":      tokens.IsAdmin,
	})
}

type loginRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (a *AuthController) Login(c *gin.Context) {
	httputils.DisableCors(c)

	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	tokens, err := a.authService.LoginEmployee(c.Request.Context(), &dto.LoginEmployeeRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		a.l.Errorf("can`t login employee: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Can`t login employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
		"isAdmin":      tokens.IsAdmin,
	})
}

type refreshTokensRequest struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	httputils.DisableCors(c)

	var req refreshTokensRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	payload, err := a.authService.VerifyEmployeeAccessToken(c.Request.Context(), &dto.VerifyEmployeeAccessTokenRequest{AccessToken: req.AccessToken})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) && !errors.Is(err, jwt.ErrTokenNotValidYet) {
		a.l.Errorf("failed to verify token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		a.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	tokens, err := a.authService.RefreshTokens(c.Request.Context(), &dto.RefreshEmployeeTokensRequest{
		InfoCardID:   infoCardID,
		RefreshToken: req.RefreshToken,
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		a.l.Warnf("expired refresh token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired refresh token"})
		return
	}
	if err != nil {
		a.l.Errorf("refresh tokens for employee: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t refresh tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
		"isAdmin":      tokens.IsAdmin,
	})
}
