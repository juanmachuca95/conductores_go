package controllers

type AppController struct {
	Auth   interface{ AuthController }
	Driver interface{ DriverController }
}
