package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const CacheLimit int = 100

var LRUCache []string

var CacheMap = map[string]string{}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// Health Check
	e.GET("/health", health)

	// Get Cache Value
	e.GET("/cache", getCache)

	e.POST("/cache", putCache)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
