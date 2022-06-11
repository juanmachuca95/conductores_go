package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/conductores_go/infrastructure/datastore"
	"github.com/juanmachuca95/conductores_go/infrastructure/router"
	"github.com/juanmachuca95/conductores_go/registry"
)

func main() {
	loadEnv()
	apiPort := os.Getenv("API_PORT")

	db := datastore.NewMySQLClient()
	re := registry.NewRegistry(db)

	app := re.NewAppController()
	mdw := re.NewMiddlewares()
	r := router.SetupRouter(app, mdw)

	r.Run(apiPort)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
