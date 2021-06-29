package go1inch

import (
	"context"
	"errors"
)

// Quote gets quote for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Quote(ctx context.Context, network, fromTokenAddress, toTokenAddress, amount string, opts *QuoteOpts) (*QuoteRes, int, error) {
	endpoint := "/quote"

	if fromTokenAddress == "" || toTokenAddress == "" || amount == "" {
		return nil, 0, errors.New("required parameter is missing")
	}

	var queries = make(map[string]interface{})

	queries["fromTokenAddress"] = fromTokenAddress
	queries["toTokenAddress"] = toTokenAddress
	queries["amount"] = amount

	if opts != nil {

		if opts.Fee != "" {
			queries["fee"] = opts.Fee
		}
		if opts.Protocols != "" {
			queries["protocols"] = opts.Protocols
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
		if opts.MainRouteParts != "" {
			queries["mainRouteParts"] = opts.MainRouteParts
		}
		if opts.Parts != "" {
			queries["parts"] = opts.Parts
		}
	}

	var dataRes QuoteRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
