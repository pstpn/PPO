package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"course/internal/service"
	"course/pkg/logger"
)

type Controller struct {
	handler *gin.Engine
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	return &Controller{handler: handler}
}

func (c *Controller) SetAuthRoute(l logger.Interface, authService service.AuthService) {
	a := &authController{
		l:           l,
		authService: authService,
	}

	c.handler.POST("/login", a.Login)
	c.handler.POST("/register", a.Register)

	// Disable CORS
	c.handler.OPTIONS("/*any", c.disableCors)
}

func (c *Controller) disableCors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
}
