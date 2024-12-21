package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"danieletortora.com/cryptomaster/config"
	"danieletortora.com/cryptomaster/models"
)

type GetRateResponseBody struct {
	Timestamp string  `json:"timestamp"`
	Low       string  `json:"low"`
	High      string  `json:"high"`
	Last      string  `json:"last"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
}

// Function variable to get the API URL, can be overridden in tests
var getCexUrl = func() string {
	return config.CexUrl
}

func GetRate(crypto, currency string) (*models.CryptoRate, error) {

	apiUrl := getCexUrl() + fmt.Sprintf("/api/ticker/%v/%v", crypto, strings.ToUpper(currency))

	res, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error with code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	responseBody := GetRateResponseBody{}

	var jsonErr = json.Unmarshal(body, &responseBody)

	if jsonErr != nil {
		return nil, jsonErr
	}

	lastPrice, err := strconv.ParseFloat(responseBody.Last, 64)
	if err != nil {
		return nil, err
	}

	return &models.CryptoRate{Currency: currency, Price: lastPrice}, nil
}
