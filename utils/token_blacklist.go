package utils

import "sync"

var tokenBlacklist = make(map[string]bool)
var mutex = &sync.Mutex{}

// Add token to blacklist
func BlacklistToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	tokenBlacklist[token] = true
}

// Check if token is blacklisted
func IsTokenBlacklisted(token string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	return tokenBlacklist[token]
}
