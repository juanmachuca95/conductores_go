package api

import (
	"github.com/gin-gonic/gin"
	serviceAuth "github.com/juanmachuca95/spaceguru/services/auth/handlers"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	/* Services */
	authSvc := serviceAuth.NewServiceHTTPAuth()

	/* Auth - Services */
	r.POST("/login", authSvc.LoginHandler)

	return r
}
