package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRate(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	defer mockServer.Close()

	// Override the getCexUrl function to return the mock server URL
	originalGetCexUrl := getCexUrl
	getCexUrl = func() string {
		return mockServer.URL
	}
	defer func() {
		getCexUrl = originalGetCexUrl
	}()

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
