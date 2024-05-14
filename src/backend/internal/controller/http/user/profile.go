package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	httputils "course/internal/controller/http/utils"
	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

const multiFormSizeDefault = 10000000

type ProfileController struct {
	l               logger.Interface
	documentService service.DocumentService
	fieldService    service.FieldService
	authService     service.AuthService
	photoService    service.PhotoService
}

func NewProfileController(
	l logger.Interface,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) *ProfileController {
	return &ProfileController{
		l:               l,
		documentService: documentService,
		fieldService:    fieldService,
		authService:     authService,
		photoService:    photoService,
	}
}

type fillProfileRequest struct {
	DocumentSerialNumber string  `json:"serialNumber"`
	DocumentType         string  `json:"documentType"`
	DocumentFields       []field `json:"documentFields"`
}

type field struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (p *ProfileController) FillProfile(c *gin.Context) {
	httputils.DisableCors(c)

	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	err = c.Request.ParseMultipartForm(multiFormSizeDefault)
	if err != nil {
		p.l.Errorf("failed to parse multipart form: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request"})
		return
	}

	file, _, err := c.Request.FormFile("profileData")
	if err != nil {
		p.l.Errorf("failed to parse profile data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		p.l.Errorf("failed to read profile data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}

	var req fillProfileRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		p.l.Errorf("failed to decode profile data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}

	document, err := p.documentService.CreateDocument(c.Request.Context(), &dto.CreateDocumentRequest{
		SerialNumber: req.DocumentSerialNumber,
		InfoCardID:   infoCardID,
		DocumentType: model.ToDocumentTypeFromString(req.DocumentType).Int(),
	})
	if err != nil {
		p.l.Errorf("failed to create document: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	for _, newField := range req.DocumentFields {
		_, err = p.fieldService.CreateDocumentField(c.Request.Context(), &dto.CreateDocumentFieldRequest{
			DocumentID: document.ID.Int(),
			Value:      newField.Value,
			Type:       model.ToFieldTypeFromString(newField.Type).Int(),
		})
		if err != nil {
			p.l.Errorf("failed to create document field: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document field"})
			return
		}
	}

	f, err := c.FormFile("image")
	if err != nil {
		p.l.Errorf("failed to get image from form file: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t get image from form file"})
		return
	}
	photo, err := f.Open()
	if err != nil {
		p.l.Errorf("failed to open form file: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Can`t open form file"})
		return
	}
	defer photo.Close()

	photoData, err := io.ReadAll(photo)
	if err != nil {
		p.l.Errorf("failed to read photo data: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t read photo data"})
		return
	}

	_, err = p.photoService.CreatePhoto(c.Request.Context(), &dto.CreatePhotoRequest{
		DocumentID: document.ID.Int(),
		Data:       photoData,
	})
	if err != nil {
		p.l.Errorf("failed to create photo: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.Status(http.StatusCreated)
}

func (p *ProfileController) GetProfile(c *gin.Context) {
	httputils.DisableCors(c)

	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
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

	documentFields, err := p.fieldService.ListDocumentFields(c.Request.Context(), &dto.ListDocumentFieldsRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		p.l.Errorf("failed to list document fields: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list document fields"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"documentType":   document.Type.String(),
		"serialNumber":   document.SerialNumber,
		"documentFields": p.modelToFields(documentFields),
	})
}

func (p *ProfileController) GetEmployeeImage(c *gin.Context) {
	httputils.DisableCors(c)

	payload, err := httputils.VerifyAccessToken(c, p.l, p.authService)
	if err != nil {
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		p.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	document, err := p.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		p.l.Errorf("failed to get document by infoCard ID: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get info card document"})
		return
	}

	photoData, err := p.photoService.GetPhoto(c.Request.Context(), &dto.GetPhotoRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		p.l.Errorf("failed to get employee photo: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee photo"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", photoData.Data)
}

func (p *ProfileController) modelToFields(documentFields []*model.Field) []field {
	fields := make([]field, 0)
	for _, documentField := range documentFields {
		fields = append(fields, field{
			Type:  documentField.Type.String(),
			Value: documentField.Value,
		})
	}
	return fields
}
