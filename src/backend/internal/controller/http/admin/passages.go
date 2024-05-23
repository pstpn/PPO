package admin

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	httputils "course/internal/controller/http/utils"
	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

type PassageController struct {
	l                 logger.Interface
	documentService   service.DocumentService
	checkpointService service.CheckpointService
	authService       service.AuthService
}

func NewPassageController(
	l logger.Interface,
	documentService service.DocumentService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) *PassageController {
	return &PassageController{
		l:                 l,
		documentService:   documentService,
		checkpointService: checkpointService,
		authService:       authService,
	}
}

type createPassageRequest struct {
	InfoCardID   int64     `json:"infoCardID"`
	DocumentType string    `json:"documentType"`
	Time         time.Time `json:"time"`
}

func (p *PassageController) CreatePassage(c *gin.Context) {
	httputils.DisableCors(c)

	_, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}

	var req createPassageRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		p.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: req.InfoCardID,
	})
	if err != nil {
		p.l.Errorf("failed to get document by infoCard ID: %s", err.Error())

		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
		}
		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
		return
	}

	_, err = p.checkpointService.CreatePassage(c.Request.Context(), &dto.CreatePassageRequest{
		CheckpointID: 1,
		DocumentID:   document.ID.Int(),
		Type:         model.ToDocumentTypeFromString(req.DocumentType).Int(),
		Time:         &req.Time,
	})
	if err != nil {
		p.l.Errorf("failed to create passage: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create passage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
