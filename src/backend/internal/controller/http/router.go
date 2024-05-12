package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

// SetInfoCardRoute TODO: this is admin route
func (c *Controller) SetInfoCardRoute(l logger.Interface, infoCardService service.InfoCardService, authService service.AuthService) {
	//i := admin.NewInfoCardController(l, infoCardService, authService)

	//c.handler.POST("/infocards", i.CreateInfoCard)
}

func (c *Controller) SetProfileRoute(
	l logger.Interface,
	documentService service.DocumentService,
	fieldService service.FieldService,
	authService service.AuthService,
	photoService service.PhotoService,
) {
	p := user.NewProfileController(l, documentService, fieldService, authService, photoService)

	c.handler.POST("/profile", p.FillProfile)
}
