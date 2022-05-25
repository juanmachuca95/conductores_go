package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	"github.com/juanmachuca95/spaceguru/internal/middleware"

	serviceAuth "github.com/juanmachuca95/spaceguru/services/auth/handlers"
	serviceCond "github.com/juanmachuca95/spaceguru/services/conductores/handlers"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies(nil)

	/* Conección */
	db := database.NewMySQLClient()

	/* Seeder DB*/
	/*
		seeder := seed.NewSeeder(db)
		seeder.UsersSeed()         // siembra usuarios in db
		seeder.RolesSeed()         // sembrar roles
		seeder.RolesUsersSeeder()  // sembrar rolesuser
		seeder.ConductoresSeeder() // sembrar conductores
	*/

	/* Services */
	authSvc := serviceAuth.NewServiceHTTPAuth(db)
	condSvc := serviceCond.NewServiceHTTPConductores(db)

	/* Auth - Services */
	r.POST("/login", authSvc.LoginHandler)
	r.POST("/register", authSvc.RegisterHandler)

	/* Conductores - Services */
	r.Use(middleware.AuthorizeJWT())
	r.GET("/conductores", condSvc.GetConductoresHandler)
	r.GET("/conductoresdisponibles", condSvc.GetConductoresDisponiblesHandler)

	return r
}
