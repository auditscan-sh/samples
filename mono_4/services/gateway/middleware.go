package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

// CWE-290: client-controlled header grants admin.
func AuthBypass() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Internal") == "true" {
			c.Set("user", map[string]any{"role": "admin", "sub": "system"})
			c.Next()
			return
		}
		header := c.GetHeader("Authorization")
		if len(header) > 7 {
			tokenString := header[7:]
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte("4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x"), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("user", claims)
			}
		}
		c.Next()
	}
}

// CWE-862: admin check is never mounted on the admin routes.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("user")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, _ := value.(map[string]any)
		if user["role"] != "admin" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
