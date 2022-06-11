package repository

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type mdwRepository struct {
	jwtRepository repository.JwtRepository
}

func NewMiddlewareRepository(j repository.JwtRepository) repository.MiddlewareRepository {
	return &mdwRepository{jwtRepository: j}
}

func (mdw *mdwRepository) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(403)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := mdw.jwtRepository.ValidateToken(tokenString)
		if token != nil && token.Valid {
			roles, err := mdw.jwtRepository.ExtractDataInfoFromJWT(tokenString)
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
