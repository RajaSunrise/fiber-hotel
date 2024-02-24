package middleware

import (
	"crypto/rand"
	"log"
)

var JwtSecret []byte


func generateJWTSecret() {
	
	JwtSecret = make([]byte, 64)
	_, err := rand.Read(JwtSecret)
	if err != nil {
		log.Fatalf("Failed to generate JWT secret: %v", err)
	}
}


func init() {
	generateJWTSecret()
}
