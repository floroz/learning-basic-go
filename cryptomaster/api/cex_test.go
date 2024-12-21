package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mockServer *httptest.Server
var originalGetCexUrl func() string

func TestMain(m *testing.M) {
	// Setup mock server
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "timestamp": "1638316800",
            "low": "50000.00",
            "high": "60000.00",
            "last": "55000.00",
            "ask": 56000.00,
            "bid": 54000.00
        }`))
	}))

	// Save the original getCexUrl function
	originalGetCexUrl = getCexUrl

	// Run tests
	code := m.Run()

	// Teardown mock server
	mockServer.Close()

	// Restore the original getCexUrl function
	getCexUrl = originalGetCexUrl

	// Exit with the test code
	os.Exit(code)
}

func setup() {
	// Override the getCexUrl function to return the mock server URL
	getCexUrl = func() string {
		return mockServer.URL
	}
}

func teardown() {
	// Restore the original getCexUrl function
	getCexUrl = originalGetCexUrl
}

func TestGetRate(t *testing.T) {
	setup()
	defer teardown()

	// Call the function with the mock server URL
	cryptoRate, err := GetRate("BTC", "USD")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedPrice := 55000.00
	if cryptoRate.Price != expectedPrice {
		t.Errorf("Expected price %v, got %v", expectedPrice, cryptoRate.Price)
	}
}
