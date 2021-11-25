package main

import (
	"log"
	_rand "math/rand"
	"os"
	"strconv"

	"github.com/dblclik/EasyCache/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DefaultCacheLimit int = 10

var (
	CacheLimit     int
	LRUCache       *utils.DoublyLinkedList = utils.InitDoublyList()
	CacheMap                               = map[string]string{}
	InstanceID     uint32                  = _rand.Uint32()
	HashCheckinURL string
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
func initDotEnv(envFile string) bool {
	// load .env file
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env file")
		return false
	}

	return true
}

// return the value of the key
func goDotEnvVariable(key string) string {
	log.Println(os.Getenv(key))
	return os.Getenv(key)
}

func main() {
	// initialize the env file
	envLoaded := initDotEnv(".env")

	if !envLoaded {
		panic("env file could not be loaded; panicking!")
	}

	// before coming online, cache node needs to check in
	// ** IF CHECKIN_HOST ENV IS "" THEN WILL NOT CHECK **
	HashCheckinURL = goDotEnvVariable("CHECKIN_HOST")

	if HashCheckinURL != "" {
		// initiate BLOCKING checkin against CHECKIN_HOST
		checkinOK := checkinToHashRing(HashCheckinURL)
		if !checkinOK {
			log.Fatalln("ERROR: Was not able to perform checkin, aborting now!")
		}
	} else {
		log.Println("No checkin host provided, will not assume distributed")
	}

	// Initialize Cache Limit
	log.Println("Initial Cache Limit set to: ", CacheLimit)
	godotenv.Load()
	var err error
	CacheLimit, err = strconv.Atoi(goDotEnvVariable("CACHE_SIZE_LIMIT"))
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
