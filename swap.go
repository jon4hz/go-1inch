package go1inch

import (
	"context"
	"errors"
)

// Swap gets swap for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Swap(ctx context.Context, network, fromTokenAddress, toTokenAddress, amount, fromAddress string, slippage int64, opts *SwapOpts) (*SwapRes, int, error) {
	endpoint := "/swap"

	if fromTokenAddress == "" || toTokenAddress == "" || amount == "" || fromAddress == "" {
		return nil, 0, errors.New("required parameter is missing")
	}

	var queries = make(map[string]interface{})

	queries["fromTokenAddress"] = fromTokenAddress
	queries["toTokenAddress"] = toTokenAddress
	queries["amount"] = amount
	queries["fromAddress"] = fromAddress
	queries["slippage"] = slippage

	if opts != nil {
		queries["burnChi"] = opts.BurnChi
		queries["allowPartialFill"] = opts.AllowPartialFill
		queries["disableEstimate"] = opts.DisableEstimate

		if opts.Protocols != "" {
			queries["protocols"] = opts.Protocols
		}
		if opts.DestReceiver != "" {
			queries["destReceiver"] = opts.DestReceiver
		}
		if opts.ReferrerAddress != "" {
			queries["referrerAddress"] = opts.ReferrerAddress
		}
		if opts.Fee != "" {
			queries["fee"] = opts.Fee
		}
		if opts.GasPrice != "" {
			queries["gasPrice"] = opts.GasPrice
		}
		if opts.ComplexityLevel != "" {
			queries["complexityLevel"] = opts.ComplexityLevel
		}
		if opts.ConnectorTokens != "" {
			queries["connectorTokens"] = opts.ConnectorTokens
		}
		if opts.GasLimit != "" {
			queries["gasLimit"] = opts.GasLimit
		}
		if opts.Parts != "" {
			queries["parts"] = opts.Parts
		}
		if opts.VirtualParts != "" {
			queries["virtualParts"] = opts.VirtualParts
		}
		if opts.MainRouteParts != "" {
			queries["mainRouteParts"] = opts.MainRouteParts
		}
	}

	var dataRes SwapRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
