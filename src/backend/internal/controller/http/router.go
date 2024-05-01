package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

type Controller struct {
	handler *gin.Engine
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Adding healthcheck
	handler.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Disable CORS
	handler.OPTIONS("/*any", disableCors)

	return &Controller{handler: handler}
}

func (c *Controller) SetAuthRoute(l logger.Interface, authService service.AuthService) {
	a := &authController{
		l:           l,
		authService: authService,
	}

	c.handler.POST("/login", a.Login)
	c.handler.POST("/register", a.Register)
	c.handler.POST("/refresh", a.RefreshTokens)
}

func (c *Controller) SetInfoCardRoute(l logger.Interface, infoCardService service.InfoCardService, authService service.AuthService) {
	i := &infoCardController{
		l:               l,
		infoCardService: infoCardService,
		authService:     authService,
	}

	c.handler.POST("/infocards", i.CreateInfoCard)
}

func disableCors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
}

func parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", fmt.Errorf("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", fmt.Errorf("token is empty")
	}

	return headerParts[1], nil
}

func verifyAccessToken(c *gin.Context, l logger.Interface, authService service.AuthService) (string, error) {
	accessToken, err := parseAuthHeader(c)
	if err != nil {
		l.Errorf("failed to parse auth header: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid auth header"})
		return "", err
	}

	id, err := authService.VerifyEmployeeAccessToken(c.Request.Context(), &dto.VerifyEmployeeAccessTokenRequest{AccessToken: accessToken})
	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		l.Warnf("expired token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
		return "", err
	} else if err != nil {
		l.Errorf("failed to verify token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return "", err
	}

	return id, nil
}
