package admin

import (
	"net/http"

	httputils "course/internal/controller/http/utils"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
	"github.com/gin-gonic/gin"
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

func (i *InfoCardController) ListInfoCards(c *gin.Context) {
	httputils.DisableCors(c)

	payload, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}
	_, err = payload.GetInfoCardID()
	if err != nil {
		i.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	fullInfoCards, err := i.infoCardService.ListInfoCards(c.Request.Context(), &dto.ListInfoCardsRequest{
		Pagination: &postgres.Pagination{
			PageNumber: -1,
			PageSize:   -1,
			Filter: postgres.FilterOptions{
				Pattern: c.Query("pattern"),
				Column:  c.Query("field"),
			},
			Sort: postgres.SortOptions{
				Direction: postgres.SortDirectionFromString(c.Query("sort")),
				Columns:   []string{c.Query("field")},
			},
		},
	})
	if err != nil {
		i.l.Errorf("failed to list fullInfoCards: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list info cards"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"infoCards": fullInfoCards,
	})
}
