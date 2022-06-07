package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/juanmachuca95/conductores_go/internal/service_jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(403)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			roles, err := service.JWTAuthService().ExtractDataInfoFromJWT(tokenString)
			if err != nil {
				c.AbortWithStatus(403)
			}

			next := false
			for _, role := range roles.([]interface{}) {
				role := fmt.Sprintf("%v", role)
				if role == "admin" {
					next = true
				}
			}
			if next {
				c.Next()
			} else {
				c.AbortWithStatus(403)
			}
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
