package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/juanmachuca95/conductores_go/internal/databases"
	"github.com/juanmachuca95/conductores_go/internal/middleware"

	serviceAuth "github.com/juanmachuca95/conductores_go/services/auth/handlers"
	serviceCond "github.com/juanmachuca95/conductores_go/services/conductores/handlers"
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

	/* Conductores - Services */
	r.Use(middleware.AuthorizeJWT())
	r.POST("/createadmin", authSvc.RegisterHandler)
	r.GET("/conductores", condSvc.GetConductoresHandler)
	r.POST("/createconductor", condSvc.CreateConductorHandler)
	r.GET("/conductoresdisponibles", condSvc.GetConductoresDisponiblesHandler)

	return r
}
