package private

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-iron-session/src/lib/auth"
)

func DevRouterProtected(r *gin.RouterGroup) {
	devGroup := r.Group("/dev")

	devGroup.GET("/protected", func(c *gin.Context) {
		user, exists := auth.GetUserFromContext(c)
		if !exists {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(200, gin.H{
			"message": "Protected route",
			"user":    user,
		})
	})

	devGroup.GET("/me", func(c *gin.Context) {
		user, _ := auth.GetUserFromContext(c)
		c.JSON(200, gin.H{"user": user})
	})
}