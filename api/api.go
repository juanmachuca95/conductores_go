package api

func Start(port string) {
	r := InitRoute()
	r.Run(port)
}
