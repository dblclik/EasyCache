package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dblclik/EasyCache/models"
)

const maxRetries = 4

var (
	backoffTime int
	timesTried  int
)

func httpCheckin(checkinURL string) (body string, err error) {
	requestBody := url.Values{}
	requestBody.Add("instance", string(InstanceID))
	resp, err := http.PostForm(checkinURL, requestBody)
	if err != nil {
		return "", fmt.Errorf("CheckinError: Could not checkin to Hash Service")
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("CheckinError: Could not read response from Hash Service")
	}
	return body, nil
}

func checkinToHashRing(checkinURL string, checkinResponse *models.HashResponse) bool {
	backoffTime = 2
	timesTried = 0

	for {
		if timesTried > maxRetries {
			break
		}
		responseBody, err := httpCheckin(checkinURL)
		if err != nil {
			log.Println("Response from Checkin Service Was: ", responseBody)
			time.Sleep(time.Duration(backoffTime) * time.Second)
			timesTried++
			backoffTime *= 2
		}

	}

	return false
}
