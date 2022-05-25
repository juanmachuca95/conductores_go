package conductores

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	conductores "github.com/juanmachuca95/spaceguru/services/conductores/gateway"
)

type ServiceHTTPConductores struct {
	conductores.ConductorGateway
}

func NewServiceHTTPConductores(db *sql.DB) *ServiceHTTPConductores {
	return &ServiceHTTPConductores{conductores.NewConductorGateway(db)}
}

func (s *ServiceHTTPConductores) GetConductoresHandler(c *gin.Context) {
	var request struct {
		Page int `json:"page"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(request.Page)
	conductores, err := s.GetConductores(request.Page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "conductores": conductores})
}
