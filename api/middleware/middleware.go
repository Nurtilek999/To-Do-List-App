package middleware

import (
	"Pet_1/pkg/jwttoken"
	"Pet_1/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.ValidateJWT(c.Request)
		if err != nil {
			response.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
