package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/spaceguru/api"
)

func main() {
	loadEnv()
	apiPort := os.Getenv("API_PORT")
	api.Start(apiPort)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
