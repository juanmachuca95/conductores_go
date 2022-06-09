package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/conductores_go/usescases/interactor"
)

type driverController struct {
	driverInteractor interactor.DriverInteractor
}

type DriverController interface {
	GetDriversWithPagination(*gin.Context)
}

func NewDriverController(i interactor.DriverInteractor) DriverController {
	return &driverController{driverInteractor: i}
}

func (d *driverController) GetDriversWithPagination(c *gin.Context) {
	var request struct {
		Pager int64 `json:"pager"`
	}
	if err := c.ShouldBindJSON(&request); !errors.Is(err, nil) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	drivers, err := d.driverInteractor.GetDriversWithPagination(request.Pager)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, drivers)
	return
}
