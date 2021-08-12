package go1inch_test

import (
	"context"
	"testing"

	go1inch "github.com/jon4hz/go-1inch"
)

func TestHealthCheckReponse(t *testing.T) {
	client := go1inch.NewClient()
	res, _, err := client.Healthcheck(context.Background(), "eth")
	if err != nil {
		t.Error(err)
	}
	if res.Status != "OK" {
		t.Errorf("healthcheck returned %s, expected OK", res.Status)
	}
	t.Log(res)
}

func TestQuoteWithoutOpts(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.Quote(context.Background(), "eth",
		"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"100000000000000000000",
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestQuoteWithOpts(t *testing.T) {
	client := go1inch.NewClient()
	res, _, err := client.Quote(context.Background(), "eth",
		"0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		"0x111111111117dc0aa78b770fa6a738034120c302",
		"100000000000000000000",
		&go1inch.QuoteOpts{
			Fee:             "1",
			GasPrice:        "1",
			ComplexityLevel: "2",
			GasLimit:        "10000000",
			VirtualParts:    "50",
			Parts:           "50",
			MainRouteParts:  "10",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestQuoteWithMissingParameter(t *testing.T) {
	client := go1inch.NewClient()

	_, _, err := client.Quote(context.Background(), "eth",
		"",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"100000000000000000000",
		nil,
	)
	if err == nil {
		t.Error("getting swap should have failed. fromTokenAddress is missing.")
	}
}

func TestSwapWithoutOpts(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.Swap(context.Background(), "eth",
		"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"100000000000000000000",
		"0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5",
		1,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestSwapWithOpts(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.Swap(context.Background(), "eth",
		"0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		"0x111111111117dc0aa78b770fa6a738034120c302",
		"100000000000000000000",
		"0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5",
		1,
		&go1inch.SwapOpts{
			DestReceiver:     "0x52bc44d5378309ee2abf1539bf71de1b7d7be3b5",
			ReferrerAddress:  "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
			Fee:              "1",
			GasPrice:         "1",
			BurnChi:          false,
			ComplexityLevel:  "2",
			AllowPartialFill: true,
			DisableEstimate:  true,
			GasLimit:         "10000000",
			Parts:            "50",
			VirtualParts:     "50",
			MainRouteParts:   "10",
		})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestSwapWithMissingParameter(t *testing.T) {
	client := go1inch.NewClient()

	_, _, err := client.Swap(context.Background(), "eth",
		"",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"100000000000000000000",
		"",
		0,
		nil,
	)
	if err == nil {
		t.Error("getting swap should have failed. fromTokenAddress is missing.")
	}
}

func TestGetProtocols(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.Protocols(context.Background(), "matic")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetProtocolsImages(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.ProtocolsImages(context.Background(), "bsc")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetTokens(t *testing.T) {
	client := go1inch.NewClient()

	res, _, err := client.Tokens(context.Background(), "matic")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestApproveSpender(t *testing.T) {
	client := go1inch.NewClient()
	res, _, err := client.ApproveSpender(context.Background(), "bsc")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestApproveCalldataWithoutOpts(t *testing.T) {
	client := go1inch.NewClient()
	res, _, err := client.ApproveCalldata(
		context.Background(),
		"eth",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestApproveCalldataWithOpts(t *testing.T) {
	client := go1inch.NewClient()
	res, _, err := client.ApproveCalldata(
		context.Background(),
		"eth",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		&go1inch.ApproveCalldataOpts{
			Amount: "100000",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestWBTC(t *testing.T) {
	client := go1inch.NewClient()
	_, code, err := client.Quote(context.Background(), "matic", "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", "0x2791bca1f2de4661ed88a30c99a7a9449aa84174", "100000000", nil)
	if err != nil {
		t.Error(err)
	}
	if code != 200 {
		t.Error("expected 200, got ", code)
	}
}
