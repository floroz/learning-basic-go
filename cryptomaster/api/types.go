package api

type GetRateResponseBody struct {
	Timestamp string  `json:"timestamp"`
	Low       string  `json:"low"`
	High      string  `json:"high"`
	Last      string  `json:"last"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
}
