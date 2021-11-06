package main

import (
	"github.com/dblclik/EasyCache/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const CacheLimit int = 100

// TODO: Need to switch to doubly linked list
var LRUCache *utils.DoublyLinkedList = utils.InitDoublyList()

var CacheMap = map[string]string{}

/* TODO:
- Implement DLL + Updating (DONE)
- Add tests for existing functions (DONE)
- Consider adding auth module
- Add env file for config variables
- Add Config Vars to HEALTH
- Enable resource updates on CONFIG vars (that can be changed)
- Start time tracking for all actions + logging these
*/

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

	// Put  (key, value) into cache
	e.POST("/cache", putCache)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
