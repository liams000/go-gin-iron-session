package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		isAdmin := authHeader == os.Getenv("ADMIN_KEY")

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}