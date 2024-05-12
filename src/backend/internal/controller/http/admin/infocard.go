package admin

import (
	"course/internal/service"
	"course/pkg/logger"
)

type InfoCardController struct {
	l               logger.Interface
	infoCardService service.InfoCardService
	authService     service.AuthService
}

func NewInfoCardController(l logger.Interface, infoCardService service.InfoCardService, authService service.AuthService) *InfoCardController {
	return &InfoCardController{
		l:               l,
		infoCardService: infoCardService,
		authService:     authService,
	}
}
