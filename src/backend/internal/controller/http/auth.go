package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
)

type authController struct {
	l           logger.Interface
	authService service.AuthService
}

type registerRequest struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	CompanyID   int64  `json:"companyID"`
	Post        string `json:"post"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (a *authController) Register(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse("02.01.2006", req.DateOfBirth)
	if err != nil {
		a.l.Errorf("incorrect request body date of birth: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	employee, err := a.authService.RegisterEmployee(c.Request.Context(), &dto.RegisterEmployeeRequest{
		PhoneNumber: req.PhoneNumber,
		FullName:    req.Name + " " + req.Surname,
		CompanyID:   req.CompanyID,
		Post:        model.ToPostTypeFromString(req.Post).Int(),
		Password: &model.Password{
			Value:    req.Password,
			IsHashed: false,
		},
		DateOfBirth: &birthDate,
	})
	if err != nil {
		a.l.Errorf("can`t register employee: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": employee.Password})
}

type loginRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (a *authController) Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := a.authService.LoginEmployee(c.Request.Context(), &dto.LoginEmployeeRequest{
		PhoneNumber: req.PhoneNumber,
		Password: &model.Password{
			Value:    req.Password,
			IsHashed: false,
		},
	})
	if err != nil {
		a.l.Errorf("can`t login employee: %s", err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
