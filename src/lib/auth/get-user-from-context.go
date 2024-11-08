package auth

import "github.com/gin-gonic/gin"

// GetUserFromContext retrieves user data from Gin context
func GetUserFromContext(c *gin.Context) (*TokenData, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	userData, ok := user.(*TokenData)
	return userData, ok
}