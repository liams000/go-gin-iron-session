package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

const (
	Param = "Fe_"
)

// GenerateToken creates a new auth token
func (m *Manager) GenerateToken(userID, username string) (string, error) {
	token := TokenData{
		UserID:    userID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(m.config.TokenExpiration),
	}

	payload, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("failed to marshal token: %w", err)
	}

	h := hmac.New(sha256.New, []byte(m.config.SecretKey))
	h.Write(payload)
	signature := h.Sum(nil)

	tokenStr := fmt.Sprintf("%s.%s",
		base64.URLEncoding.EncodeToString(payload),
		base64.URLEncoding.EncodeToString(signature))

	return Param + tokenStr, nil
}