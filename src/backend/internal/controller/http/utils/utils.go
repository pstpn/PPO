package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

func DisableCors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
}

func VerifyAccessToken(c *gin.Context, l logger.Interface, authService service.AuthService) (*model.Payload, error) {
	accessToken, err := parseAuthHeader(c)
	if err != nil {
		l.Errorf("failed to parse auth header: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid auth header"})
		return nil, err
	}

	payload, err := authService.VerifyEmployeeAccessToken(c.Request.Context(), &dto.VerifyEmployeeAccessTokenRequest{AccessToken: accessToken})
	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		l.Warnf("expired token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
		return nil, err
	} else if err != nil {
		l.Errorf("failed to verify token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return nil, err
	}

	return payload, nil
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

type Field struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func ModelToFields(documentFields []*model.Field) []Field {
	fields := make([]Field, 0)
	for _, documentField := range documentFields {
		fields = append(fields, Field{
			Type:  documentField.Type.String(),
			Value: documentField.Value,
		})
	}
	return fields
}

type Passage struct {
	Type string `json:"type"`
	Time string `json:"time"`
}

func ModelToPassages(passages []*model.Passage) []Passage {
	p := make([]Passage, 0)
	for _, passage := range passages {
		p = append(p, Passage{
			Type: passage.Type.String(),
			Time: passage.Time.Add(3 * time.Hour).Format("15:04:05 (02.01.2006)"),
		})
	}
	return p
}
