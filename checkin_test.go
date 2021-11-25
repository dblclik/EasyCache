package main

import (
	"testing"
)

/* To test httpCheckin, we will check
https://postman-echo.com/status/200 -- PASS
https://postman-echo.com/status/201 -- FAIL
https://postman-echo.com/status/401 -- FAIL
*/

func Test_httpCheckin(t *testing.T) {
	passURL := "https://postman-echo.com/status/200"
	failingURL01 := "https://postman-echo.com/status/201"
	failingURL02 := "https://postman-echo.com/status/401"

	resp01, err01 := httpCheckin(passURL)
	if resp01 != true || err01 != nil {
		t.Fatal("Did not return true when given a passing (200) URL")
	}

	resp02, err02 := httpCheckin(failingURL01)
	if resp02 != false || err02 == nil {
		t.Fatalf("Expected false when given a bad, but successful (201) URL, got %v", resp02)
	}

	resp03, err03 := httpCheckin(failingURL02)
	if resp03 != false || err03 == nil {
		t.Fatalf("Expected false when given a bad (401) URL, got %v", resp03)
	}
}

func Test_checkinToHashRing(t *testing.T) {
	passURL := "https://postman-echo.com/status/200"
	failingURL02 := "https://postman-echo.com/status/401"

	resp01 := checkinToHashRing(passURL)
	if resp01 != true {
		t.Fatal("Did not return true when given a passing (200) URL")
	}

	resp03 := checkinToHashRing(failingURL02)
	if resp03 != false {
		t.Fatalf("Expected false when given a bad (401) URL, got %v", resp03)
	}
}
