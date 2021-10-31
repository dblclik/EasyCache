package main

import (
	"net/http"
	"time"

	"github.com/dblclik/EasyCache/models"
	"github.com/labstack/echo/v4"
)

// Get Cache Entry (If Exists)
func getCache(c echo.Context) error {
	key := c.QueryParam("key")
	if val, exists := CacheMap[key]; exists {
		r := &models.CacheResponse{
			Exists:    true,
			Key:       key,
			Value:     val,
			Timestamp: time.Now().Format("20060102150405"),
		}
		return c.JSON(http.StatusOK, r)
	}
	r := &models.CacheResponse{
		Exists:    false,
		Key:       key,
		Value:     "",
		Timestamp: time.Now().Format("20060102150405"),
	}
	return c.JSON(http.StatusOK, r)
}

// Put Cache Entry, Evict Key If Needed
func putCache(c echo.Context) error {
	body := new(models.CacheLoad)
	if err := c.Bind(body); err != nil {
		return err
	}

	// if entry is new, prepend to head of LRUCache, else move item to front
	if val, exists := CacheMap[body.Key]; exists {

	} else {
		// worried about this, might be inefficient to unpack LRUCache
		LRUCache = append([]string{body.Key}, LRUCache...)
	}

	if len(CacheMap) > CacheLimit {
		go evict()
	}
	r := &models.CacheLoadConfirmation{
		Key:       body.Key,
		Timestamp: time.Now().Format("20060102150405"),
	}
	return c.JSON(http.StatusOK, r)
}
