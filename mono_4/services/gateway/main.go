package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	router := gin.Default()

	// CWE-942: wildcard origin with credentials.
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
	}))

	// No security headers, no rate limit, no request size limit.
	router.GET("/api/proxy", ProxyHandler)
	router.GET("/api/fetch", DirectFetchHandler)
	router.GET("/redirect", RedirectHandler)
	router.POST("/api/upload", UploadHandler)
	router.GET("/api/env", EnvHandler)
	router.GET("/ws", WebSocketHandler)
	router.POST("/api/checkout", CheckoutHandler)
	router.GET("/api/orders", OrdersHandler)
	router.POST("/api/forward", ForwardHandler)
	router.GET("/debug/pprof", DebugHandler)
	router.GET("/debug/config", ConfigHandler)

	router.Run("0.0.0.0:8080")
}
