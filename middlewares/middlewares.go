package middlewares

import (
	"crud-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// middleware function to check the token
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
