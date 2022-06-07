package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/conductores_go/interface/controllers"
)

func SetupRouter(app controllers.AppController) *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)

	r.POST("/login", func(ctx *gin.Context) { app.Auth.Login(ctx) })

	return r
}