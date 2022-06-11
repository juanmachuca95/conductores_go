package repository

import "github.com/gin-gonic/gin"

type MiddlewareRepository interface {
	AuthorizeJWT() gin.HandlerFunc
}
