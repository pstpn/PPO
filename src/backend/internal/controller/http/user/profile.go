package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	httputils "course/internal/controller/http/utils"
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
	DocumentSerialNumber string   `json:"serialNumber"`
	DocumentType         string   `json:"documentType"`
	DocumentFields       []fields `json:"documentFields"`
}

type fields struct {
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
		// FIXME: field
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
			// FIXME: field
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

	var photoData []byte
	_, err = photo.Read(photoData)
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
