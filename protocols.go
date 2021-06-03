package go1inch

import (
	"context"
)

// Get array of all supported liquidity protocols
func (c *Client) Protocols(ctx context.Context, network string) (*ProtocolsRes, int, error) {
	endpoint := "/protocols"
	var dataRes ProtocolsRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// Get names and images of all supported protocols
func (c *Client) ProtocolsImages(ctx context.Context, network string) (*ProtocolsImagesRes, int, error) {
	endpoint := "/protocols/images"
	var dataRes ProtocolsImagesRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
