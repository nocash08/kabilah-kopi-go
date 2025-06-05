package helper

import (
	"sync"
	"time"
)

var (
	invalidatedTokens = make(map[string]time.Time)
	tokenMutex        sync.RWMutex
)

// AddToInvalidatedTokens adds a token to the invalidated tokens map
func AddToInvalidatedTokens(token string, expiresAt time.Time) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	invalidatedTokens[token] = expiresAt
}

// IsTokenInvalidated checks if a token is in the invalidated tokens map
func IsTokenInvalidated(token string) bool {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()

	if expiry, exists := invalidatedTokens[token]; exists {
		// If token has expired, remove it from the map
		if time.Now().After(expiry) {
			go cleanupToken(token)
			return false
		}
		return true
	}
	return false
}

// cleanupToken removes a token from the invalidated tokens map
func cleanupToken(token string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	delete(invalidatedTokens, token)
}

// InitTokenCleanup starts a goroutine to periodically cleanup expired tokens
func InitTokenCleanup() {
	go func() {
		for {
			time.Sleep(1 * time.Hour) // Cleanup every hour
			cleanupExpiredTokens()
		}
	}()
}

// cleanupExpiredTokens removes all expired tokens from the map
func cleanupExpiredTokens() {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	now := time.Now()
	for token, expiry := range invalidatedTokens {
		if now.After(expiry) {
			delete(invalidatedTokens, token)
		}
	}
}
