package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// VerifyToken validates a token and returns its data
func (m *Manager) VerifyToken(tokenStr string) (*TokenData, error) {
	if len(tokenStr) < 3 || tokenStr[:3] != Param {
		return nil, fmt.Errorf("invalid token format")
	}
	tokenStr = tokenStr[3:]

	parts := strings.Split(tokenStr, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid token format")
	}

	payload, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid payload encoding: %w", err)
	}

	h := hmac.New(sha256.New, []byte(m.config.SecretKey))
	h.Write(payload)
	expectedSig := h.Sum(nil)

	actualSig, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid signature encoding: %w", err)
	}

	if !hmac.Equal(expectedSig, actualSig) {
		return nil, fmt.Errorf("invalid signature")
	}

	var token TokenData
	if err := json.Unmarshal(payload, &token); err != nil {
		return nil, fmt.Errorf("invalid token data: %w", err)
	}

	if time.Now().After(token.ExpiresAt) {
		return nil, fmt.Errorf("token expired")
	}

	return &token, nil
}