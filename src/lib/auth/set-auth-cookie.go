package auth

import "github.com/gin-gonic/gin"

// SetAuthCookie sets the auth token in a cookie
func (m *Manager) SetAuthCookie(c *gin.Context, token string) {
	c.SetCookie(
		m.config.CookieName,
		token,
		int(m.config.TokenExpiration.Seconds()),
		"/",
		m.config.CookieDomain,
		m.config.CookieSecure,
		m.config.CookieHTTPOnly,
	)
}