package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

type infoCardController struct {
	l               logger.Interface
	infoCardService service.InfoCardService
	authService     service.AuthService
}

func (i *infoCardController) CreateInfoCard(c *gin.Context) {
	disableCors(c)

	phoneNumber, err := verifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	now := time.Now()
	infoCard, err := i.infoCardService.CreateInfoCard(c.Request.Context(), &dto.CreateInfoCardRequest{
		EmployeePhoneNumber: phoneNumber,
		IsConfirmed:         false,
		CreatedDate:         &now,
	})
	if err != nil {
		i.l.Errorf("can`t create infocard: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Can`t create infocard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isConfirmed": infoCard.IsConfirmed,
		"createdDate": infoCard.CreatedDate,
	})
}
