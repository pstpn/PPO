package admin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	httputils "course/internal/controller/http/utils"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
	"course/pkg/storage/postgres"
)

type InfoCardController struct {
	l                 logger.Interface
	infoCardService   service.InfoCardService
	documentService   service.DocumentService
	fieldService      service.FieldService
	checkpointService service.CheckpointService
	photoService      service.PhotoService
	authService       service.AuthService
}

func NewInfoCardController(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) *InfoCardController {
	return &InfoCardController{
		l:                 l,
		infoCardService:   infoCardService,
		documentService:   documentService,
		fieldService:      fieldService,
		checkpointService: checkpointService,
		photoService:      photoService,
		authService:       authService,
	}
}

func (i *InfoCardController) ListFullInfoCards(c *gin.Context) {
	httputils.DisableCors(c)

	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
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

func (i *InfoCardController) GetFullInfoCard(c *gin.Context) {
	httputils.DisableCors(c)

	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		i.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
		return
	}

	document, err := i.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		i.l.Errorf("failed to get document by infoCard ID: %s", err.Error())

		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
		}
		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
		return
	}

	documentFields, err := i.fieldService.ListDocumentFields(c.Request.Context(), &dto.ListDocumentFieldsRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		i.l.Errorf("failed to list document fields: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list document fields"})
		return
	}

	passages, err := i.checkpointService.ListPassages(c.Request.Context(), &dto.ListPassagesRequest{DocumentID: document.ID.Int()})
	if err != nil {
		i.l.Errorf("failed to list passages: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list passages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"document": gin.H{
			"data": gin.H{
				"documentType": document.Type.String(),
				"serialNumber": document.SerialNumber,
			},
			"fields": httputils.ModelToFields(documentFields),
		},
		"passages": httputils.ModelToPassages(passages),
	})
}

func (i *InfoCardController) GetEmployeeInfoCardPhoto(c *gin.Context) {
	httputils.DisableCors(c)

	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		i.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
		return
	}

	document, err := i.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		i.l.Errorf("failed to get document by infoCard ID: %s", err.Error())

		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
		}
		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
		return
	}

	photoData, err := i.photoService.GetPhoto(c.Request.Context(), &dto.GetPhotoRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		i.l.Errorf("failed to get employee infoCard photo: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee info card photo"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", photoData.Data)
}

func (i *InfoCardController) ConfirmEmployeeInfoCard(c *gin.Context) {
	httputils.DisableCors(c)

	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		i.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
		return
	}

	err = i.infoCardService.ValidateInfoCard(c.Request.Context(), &dto.ValidateInfoCardRequest{
		InfoCardID:  infoCardID,
		IsConfirmed: true,
	})
	if err != nil {
		i.l.Errorf("failed to validate infoCard: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to validate info card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
