package auth

import "github.com/gin-gonic/gin"

// ClearAuthCookie removes the auth cookie
func (m *Manager) ClearAuthCookie(c *gin.Context) {
	c.SetCookie(
		m.config.CookieName,
		"",
		-1,
		"/",
		m.config.CookieDomain,
		m.config.CookieSecure,
		m.config.CookieHTTPOnly,
	)
}