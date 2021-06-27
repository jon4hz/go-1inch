package go1inch

import (
	"context"
	"fmt"
)

// Quote gets quote for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Quote(ctx context.Context, network string, request *QuoteReq) (*QuoteRes, int, error) {
	endpoint := "/quote"

	var queries = make(map[string]interface{})

	if request.FromTokenAddress == "" || request.ToTokenAddress == "" || request.Amount == "" {
		return nil, 0, fmt.Errorf("required parameter is missing")
	}

	queries["fromTokenAddress"] = request.FromTokenAddress
	queries["toTokenAddress"] = request.ToTokenAddress
	queries["amount"] = request.Amount
	queries["fee"] = request.Fee

	if request.Protocols != "" {
		queries["protocols"] = request.Protocols
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
	if request.MainRouteParts != 0 {
		queries["mainRouteParts"] = request.MainRouteParts
	}
	if request.Parts != 0 {
		queries["parts"] = request.Parts
	}
	var dataRes QuoteRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
