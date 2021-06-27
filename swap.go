package go1inch

import (
	"context"
	"fmt"
)

// Swap gets swap for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Swap(ctx context.Context, network string, request *SwapReq) (*SwapRes, int, error) {
	endpoint := "/swap"

	var queries = make(map[string]interface{})

	if request.FromTokenAddress == "" || request.ToTokenAddress == "" || request.Amount == "" || request.FromAddress == "" {
		return nil, 0, fmt.Errorf("required parameter is missing")
	}

	queries["fromTokenAddress"] = request.FromTokenAddress
	queries["toTokenAddress"] = request.ToTokenAddress
	queries["amount"] = request.Amount
	queries["fromAddress"] = request.FromAddress
	queries["slippage"] = request.Slippage

	queries["burnChi"] = request.BurnChi
	queries["allowPartialFill"] = request.AllowPartialFill
	queries["disableEstimate"] = request.DisableEstimate

	if request.Protocols != "" {
		queries["protocols"] = request.Protocols
	}
	if request.DestReceiver != "" {
		queries["destReceiver"] = request.DestReceiver
	}
	if request.ReferrerAddress != "" {
		queries["referrerAddress"] = request.ReferrerAddress
	}
	if request.Fee != 0 {
		queries["fee"] = request.Fee
	}
	if request.GasPrice != "" {
		queries["gasPrice"] = request.GasPrice
	}
	if request.ComplexityLevel != "" {
		queries["complexityLevel"] = request.ComplexityLevel
	}
	if request.ConnectorTokens != "" {
		queries["connectorTokens"] = request.ConnectorTokens
	}
	if request.GasLimit != 0 {
		queries["gasLimit"] = request.GasLimit
	}
	if request.Parts != 0 {
		queries["parts"] = request.Parts
	}
	if request.MainRouteParts != 0 {
		queries["mainRouteParts"] = request.MainRouteParts
	}

	var dataRes SwapRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
