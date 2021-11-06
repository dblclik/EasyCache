package main

import (
	"net/http"
	"time"
	"unsafe"

	"github.com/dblclik/EasyCache/models"
	"github.com/labstack/echo/v4"
)

// Health Check Handler
func health(c echo.Context) error {
	h := &models.Health{
		Status:         "okay",
		CacheDepth:     len(CacheMap),
		CacheSize:      unsafe.Sizeof(CacheMap),
		CacheItemLimit: CacheLimit,
		Timestamp:      time.Now().Format("20060102150405"),
	}
	return c.JSON(http.StatusOK, h)
}
