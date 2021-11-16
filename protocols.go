package go1inch

import (
	"context"
)

// ProtocolsImages gets names and images of all supported protocols
func (c *Client) Protocols(ctx context.Context, network string) (*ProtocolsImagesRes, int, error) {
	endpoint := "/liquidity-sources"
	var dataRes ProtocolsImagesRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
