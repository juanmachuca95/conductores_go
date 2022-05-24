package api

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()

	/*Auth register user*/
	r.POST("")

	return r
}
