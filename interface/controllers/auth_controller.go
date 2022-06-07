package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/interactor"
)

type AuthController interface {
	Login(*gin.Context)
}

type authController struct {
	authinteractor interactor.AuthInteractor
}

func NewAuthController(i interactor.AuthInteractor) AuthController {
	return &authController{authinteractor: i}
}

func (a *authController) Login(c *gin.Context) {
	var login models.Login
	if err := c.Bind(&login); !errors.Is(err, nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authinteractor.Authentication(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
