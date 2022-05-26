package conductores

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	cond "github.com/juanmachuca95/spaceguru/services/conductores/gateway"
	m "github.com/juanmachuca95/spaceguru/services/conductores/models"
)

type ServiceHTTPConductores struct {
	cond.ConductorGateway
}

func NewServiceHTTPConductores(db *sql.DB) *ServiceHTTPConductores {
	return &ServiceHTTPConductores{cond.NewConductorGateway(db)}
}

func (s *ServiceHTTPConductores) CreateConductorHandler(c *gin.Context) {
	var createConductor m.CreateConductor
	if err := c.ShouldBindJSON(&createConductor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.CreateConductor(&createConductor)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "_token": user.Token})
}

func (s *ServiceHTTPConductores) GetConductoresHandler(c *gin.Context) {
	var request struct {
		Page int `json:"page"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conductores, err := s.GetConductores(request.Page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "Success",
		"conductores": conductores,
	})
}

func (s *ServiceHTTPConductores) GetConductoresDisponiblesHandler(c *gin.Context) {
	conductoresDisponibles, err := s.GetConductoresDisponibles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Conductores disponibles que no están realizando viajes",
		"status":      "Success",
		"conductores": conductoresDisponibles,
	})
}
