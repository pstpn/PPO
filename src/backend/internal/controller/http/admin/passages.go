package admin

//
//import (
//	"errors"
//	"net/http"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jackc/pgx/v5"
//
//	httputils "course/internal/controller/http/utils"
//	"course/internal/service"
//	"course/internal/service/dto"
//	"course/pkg/logger"
//)
//
//type PassageController struct {
//	l                 logger.Interface
//	infoCardService   service.InfoCardService
//	documentService   service.DocumentService
//	checkpointService service.CheckpointService
//	authService       service.AuthService
//}
//
//func (p *PassageController) GetPassages(c *gin.Context) {
//	httputils.DisableCors(c)
//
//	_, err := httputils.VerifyAccessToken(c, p.l, p.authService)
//	if err != nil {
//		return
//	}
//
//	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
//	if err != nil {
//		p.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
//		return
//	}
//
//	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
//		InfoCardID: infoCardID,
//	})
//	if err != nil {
//		p.l.Errorf("failed to get document by infoCard ID: %s", err.Error())
//
//		status := http.StatusInternalServerError
//		if errors.Is(err, pgx.ErrNoRows) {
//			status = http.StatusNotFound
//		}
//		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
//		return
//	}
//
//	documentFields, err := i.fieldService.ListDocumentFields(c.Request.Context(), &dto.ListDocumentFieldsRequest{
//		DocumentID: document.ID.Int(),
//	})
//	if err != nil {
//		i.l.Errorf("failed to list document fields: %s", err.Error())
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list document fields"})
//		return
//	}
//
//	passages, err := i.checkpointService.ListPassages(c.Request.Context(), &dto.ListPassagesRequest{InfoCardID: infoCardID})
//	if err != nil {
//		i.l.Errorf("failed to list passages: %s", err.Error())
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list passages"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"document": gin.H{
//			"data": gin.H{
//				"documentType": document.Type.String(),
//				"serialNumber": document.SerialNumber,
//			},
//			"fields": httputils.ModelToFields(documentFields),
//		},
//		"passages": passages,
//	})
//}
