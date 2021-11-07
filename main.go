package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dblclik/EasyCache/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DefaultCacheLimit int = 10

var (
	CacheLimit int                     = DefaultCacheLimit
	LRUCache   *utils.DoublyLinkedList = utils.InitDoublyList()
	CacheMap                           = map[string]string{}
)

/* TODO:
- Implement DLL + Updating (DONE)
- Add tests for existing functions (DONE)
- Consider adding auth module
- Add env file for config variables
- Add Config Vars to HEALTH
- Enable resource updates on CONFIG vars (that can be changed)
- Start time tracking for all actions + logging these
*/

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println(os.Getenv(key))
	return os.Getenv(key)
}

func main() {
	log.Println("Initial Cache Limit set to: ", CacheLimit)
	godotenv.Load()
	CacheLimit, err := strconv.Atoi(goDotEnvVariable("CACHE_SIZE_LIMIT"))
	log.Println("Cache Limit set to: ", CacheLimit)
	if err != nil {
		log.Println("WARNING: Could not set CacheLimit with env variable, using DEFAULT")
		CacheLimit = DefaultCacheLimit
	}
	log.Println("Cache Limit set to: ", CacheLimit)
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
