package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	auth "github.com/juanmachuca95/spaceguru/services/auth/gateway"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type ServiceHTTPAuth struct {
	gtw auth.AuthGateway
}

func NewServiceHTTPAuth() *ServiceHTTPAuth {
	return &ServiceHTTPAuth{auth.NewAuthService()}
}

func (s *ServiceHTTPAuth) LoginHandler(c *gin.Context) {
	var login m.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.gtw.Login(&login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "User no autorizado"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 1) // Una día
	claims["user"] = user
	jwtSecret := os.Getenv("TOKEN_API_AUTH")
	_token, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Estas loggeado", "_token": _token})
}
