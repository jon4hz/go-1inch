package go1inch

import (
	"context"
	"testing"
)

func TestHealthCheckReponse(t *testing.T) {
	client := NewClient()
	res, _, err := client.Healthcheck(context.Background(), "eth")
	if err != nil {
		t.Error("Error while getting a healthcheck ", err)
	}
	if res.Status != "OK" {
		t.Errorf("healthcheck returned %s, expected OK", res.Status)
	}

}

func TestSwap(t *testing.T) {
	client := NewClient()

	req := &SwapReq{
		FromTokenAddress: "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		ToTokenAddress:   "0x6b175474e89094c44da98b954eedeac495271d0f",
		Amount:           "100000000000000000000",
		FromAddress:      "0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5",
		Slippage:         1,
	}

	_, _, err := client.Swap(context.Background(), "eth", req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetProtocols(t *testing.T) {
	client := NewClient()

	_, _, err := client.Protocols(context.Background(), "matic")
	if err != nil {
		t.Error(err)
	}
}

func TestGetProtocolsImages(t *testing.T) {
	client := NewClient()

	_, _, err := client.ProtocolsImages(context.Background(), "bsc")
	if err != nil {
		t.Error(err)
	}
}
