package middleware

import "github.com/gin-gonic/gin"

// BasicAuth - Adds Basic Auth username and password to API calls
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "root",
	})
}
