package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/conductores_go/interface/controllers"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

func SetupRouter(app controllers.AppController, mdw repository.MiddlewareRepository) *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)

	// Authentication
	r.POST("/login", func(ctx *gin.Context) { app.Auth.Login(ctx) })

	// Drivers
	r.Use(mdw.AuthorizeJWT())
	r.GET("/drivers", func(ctx *gin.Context) { app.Driver.GetDriversWithPagination(ctx) })
	r.GET("/driversavailable", func(ctx *gin.Context) { app.Driver.GetDriversAvailable(ctx) })

	return r
}
