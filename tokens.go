package go1inch

import "context"

// Tokens gets an array of all supported tokens (any erc20 token can be used in a quote and swap)
func (c *Client) Tokens(ctx context.Context, network string) (*TokensRes, int, error) {
	endpoint := "/tokens"
	var dataRes TokensRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
