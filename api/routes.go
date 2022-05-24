package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/spaceguru/internal/middleware"
	serviceAuth "github.com/juanmachuca95/spaceguru/services/auth/handlers"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	/* Services */
	authSvc := serviceAuth.NewServiceHTTPAuth()

	/* Auth - Services */
	r.POST("/login", authSvc.LoginHandler)
	r.POST("/register", authSvc.RegisterHandler)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthorizeJWT())
	{
	}

	return r
}
