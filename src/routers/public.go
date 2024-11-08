package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin-iron-session/src/lib/auth"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func PublicRouterGroup(e *gin.Engine) {
	publicGroup := e.Group("/api")

	publicGroup.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid request",
			})
			return
		}

		token, err := auth.AuthManager.GenerateToken("user123", req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error generating token",
			})
			return
		}

		auth.AuthManager.SetAuthCookie(c, token)

		c.JSON(http.StatusOK, LoginResponse{Token: token})
	})

	publicGroup.POST("/logout", func(c *gin.Context) {
		auth.AuthManager.ClearAuthCookie(c)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
	})
}