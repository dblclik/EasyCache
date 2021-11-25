package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const maxRetries = 4

func httpCheckin(checkinURL string) (ok bool, err error) {
	parameterizedCheckinURL := checkinURL + "?instance=" + strconv.Itoa(int(InstanceID))
	log.Println("Going to make request to URL: ", parameterizedCheckinURL)
	req, err := http.Get(parameterizedCheckinURL)
	if err != nil {
		log.Println("Error building Request to checkinURL: ", checkinURL)
		return false, fmt.Errorf("CheckinError: Could not reach the checkin URL")
	}

	// defer request close
	defer req.Body.Close()

	if req.StatusCode == http.StatusOK {
		log.Println("Successfully Checked In!")
		return true, nil
	}

	return false, fmt.Errorf("Expected 200 response code from checkin, got: " + req.Status)
}

func checkinToHashRing(checkinURL string) bool {
	backoffTime := 2
	timesTried := 0

	for {
		if timesTried > maxRetries {
			break
		}
		checkinOK, err := httpCheckin(checkinURL)
		if err != nil {
			log.Println("ERROR: httpCheckin returned: ", checkinOK)
			time.Sleep(time.Duration(backoffTime) * time.Second)
			timesTried++
			backoffTime *= 2
		}

		// If we successfully checked in, return true
		if checkinOK {
			return true
		}
	}

	return false
}
