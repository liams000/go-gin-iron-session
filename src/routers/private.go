package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-iron-session/src/lib/auth"
	"github.com/go-gin-iron-session/src/routers/private"
)

func PrivateRouterGroup(r *gin.Engine) {
	privateGroup := r.Group("/api")

	// Setup auth middle ware on all routes in here
	privateGroup.Use(auth.AuthManager.Middleware())

	private.DevRouterProtected(privateGroup)
}