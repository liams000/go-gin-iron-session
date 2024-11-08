package auth

import "time"

// TokenData represents the contents of our auth token
type TokenData struct {
	UserID    string    `json:"uid"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

// Config holds auth configuration
type Config struct {
	SecretKey        string
	TokenExpiration  time.Duration
	CookieName       string
	CookieDomain     string
	CookieSecure     bool
	CookieHTTPOnly   bool
}

// Manager handles authentication operations
type Manager struct {
	config Config
}

var AuthManager *Manager

// NewManager creates a new auth manager
func NewManager(config Config) {
	if config.TokenExpiration == 0 {
		config.TokenExpiration = 24 * time.Hour // default 24 hours
	}
	if config.CookieName == "" {
		config.CookieName = "go-gin-iron-session_auth"
	}
	AuthManager = &Manager{config: config}

	return
}