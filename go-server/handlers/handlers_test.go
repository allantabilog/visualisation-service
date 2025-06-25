package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/allantabilog/fibonacci-service/fibonacci"
)

func TestHandlePing(t *testing.T) {
	// Create a logger that discards output to keep tests quiet
	logger := log.New(ioutil.Discard, "", 0)
	fibSvc := fibonacci.NewService(logger)
	handler := NewHandler(logger, fibSvc)

	// Create a request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler.HandlePing(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify the message
	if response.Message != "fibonacci-server is up" {
		t.Errorf("handler returned unexpected message: got %v want %v", response.Message, "fibonacci-server is up")
	}

	// Verify timestamp is recent
	now := time.Now()
	if response.Timestamp.After(now) || response.Timestamp.Before(now.Add(-time.Minute)) {
		t.Errorf("handler returned unexpected timestamp: got %v", response.Timestamp)
	}
}

func TestHandleFibonacci(t *testing.T) {
	// Create a logger that discards output to keep tests quiet
	logger := log.New(ioutil.Discard, "", 0)
	fibSvc := fibonacci.NewService(logger)
	handler := NewHandler(logger, fibSvc)

	testCases := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedValue  uint64
	}{
		{"Valid request", "n=10", http.StatusOK, 55},
		{"Missing parameter", "", http.StatusBadRequest, 0},
		{"Invalid parameter", "n=invalid", http.StatusBadRequest, 0},
		{"Negative number", "n=-10", http.StatusBadRequest, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request
			req, err := http.NewRequest("GET", "/fibonacci?"+tc.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler.HandleFibonacci(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			// For successful requests, check the response body
			if tc.expectedStatus == http.StatusOK {
				var response FibonacciResponse
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}

				// Verify the value
				if response.Value != tc.expectedValue {
					t.Errorf("handler returned unexpected value: got %v want %v", response.Value, tc.expectedValue)
				}
			}
		})
	}
}
