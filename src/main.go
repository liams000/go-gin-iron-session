package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gin-iron-session/src/lib/auth"
	"github.com/go-gin-iron-session/src/routers"
	"github.com/joho/godotenv"
)

func main() {
	startup()

	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time": time.Now(),
		})
	})

	// Public routes
	routers.PublicRouterGroup(app)

	// Private routes
	routers.PrivateRouterGroup(app)

	err := app.Run()
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func startup() {
	// Initialise .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth.NewManager(auth.Config{
		SecretKey: os.Getenv("API_SESSION_KEY_PASS"),
		TokenExpiration: 24*time.Hour,
		CookieName: "go-gin-iron-session_auth",
		CookieSecure: false, // Set to true in production
		CookieHTTPOnly: true,
	})
}