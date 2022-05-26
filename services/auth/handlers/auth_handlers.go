package auth

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/juanmachuca95/spaceguru/services/auth/gateway"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type ServiceHTTPAuth struct {
	gtw auth.AuthGateway
}

func NewServiceHTTPAuth(db *sql.DB) *ServiceHTTPAuth {
	return &ServiceHTTPAuth{auth.NewAuthGateway(db)}
}

func (s *ServiceHTTPAuth) LoginHandler(c *gin.Context) {
	var login m.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.gtw.Login(&login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "_token": user.Token})
}

func (s *ServiceHTTPAuth) RegisterHandler(c *gin.Context) {
	var register m.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.gtw.Register(&register)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Ya existe una cuenta con este usuario."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "_token": user.Token})
}
