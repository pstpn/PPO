package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"course/internal/controller/http/admin"
	"course/internal/controller/http/user"
	httputils "course/internal/controller/http/utils"
	"course/internal/service"
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
	handler.OPTIONS("/*any", httputils.DisableCors)

	return &Controller{handler: handler}
}

func (c *Controller) SetAuthRoute(l logger.Interface, authService service.AuthService) {
	a := NewAuthController(l, authService)

	c.handler.POST("/login", a.Login)
	c.handler.POST("/register", a.Register)
	c.handler.POST("/refresh", a.RefreshTokens)
}

func (c *Controller) SetInfoCardRoute(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	photoService service.PhotoService,
	authService service.AuthService,
) {
	i := admin.NewInfoCardController(
		l,
		infoCardService,
		documentService,
		fieldService,
		checkpointService,
		photoService,
		authService,
	)

	c.handler.GET("/infocards", i.ListFullInfoCards)
	c.handler.GET("/infocards/:id", i.GetFullInfoCard)
	c.handler.PUT("/infocards/:id", i.ConfirmEmployeeInfoCard)
	c.handler.GET("infocard-photos/:id", i.GetEmployeeInfoCardPhoto)
}

func (c *Controller) SetProfileRoute(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) {
	p := user.NewProfileController(l, infoCardService, documentService, fieldService, authService, photoService)

	c.handler.POST("/profile", p.FillProfile)
	c.handler.GET("/profile", p.GetProfile)
	c.handler.GET("/employee-photo", p.GetEmployeePhoto)
}

func (c *Controller) SetPassageRoute(
	l logger.Interface,
	documentService service.DocumentService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) {
	p := admin.NewPassageController(l, documentService, checkpointService, authService)

	c.handler.POST("/passages", p.CreatePassage)
}
