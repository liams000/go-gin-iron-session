package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Manager) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""

		// Check Authorization header first
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// If no Authorization header, check cookie
		if token == "" {
			var err error
			token, err = c.Cookie(m.config.CookieName)
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			}
		}

		// Verify token
		tokenData, err := m.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Set user data in context
		c.Set("user", tokenData)
		c.Next()
	}
}