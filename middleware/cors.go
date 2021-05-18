package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors settings
func Cors() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type", "Cookie"}

	if gin.Mode() == gin.ReleaseMode {
		// domain name under production
		config.AllowOrigins = []string{"http://www.example.com"}
	} else {
		// local testing
		config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	}
	config.AllowCredentials = true

	return cors.New(config)
}
