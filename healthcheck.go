package go1inch

import (
	"context"
)

// Healthcheck checks if service is able to handle requests
func (c *Client) Healthcheck(ctx context.Context, network string) (*HealthcheckRes, int, error) {
	endpoint := "/healthcheck"
	var dataRes HealthcheckRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
