package go1inch

import "net/http"

type Client struct {
	Http *http.Client
}

type HealthcheckResponse struct {
	Status string `json:"status"`
}

type ApproveCalldataResponse struct {
	To       []string `json:"to"`
	Value    string   `json:"value"`
	GasPrice string   `json:"gasPrice"`
	Data     string   `json:"data"`
}
