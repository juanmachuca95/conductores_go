package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	"github.com/juanmachuca95/spaceguru/internal/middleware"
	serviceAuth "github.com/juanmachuca95/spaceguru/services/auth/handlers"
	serviceCond "github.com/juanmachuca95/spaceguru/services/conductores/handlers"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
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

	authorized := r.Group("/")
	authorized.Use(middleware.AuthorizeJWT())
	{
		/* Conductores - Services */
		r.GET("/conductores", condSvc.GetConductoresHandler)
	}

	return r
}
