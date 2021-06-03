package go1inch

import (
	"context"
	"fmt"
)

// WIP
// Get quote
func Quote(ctx context.Context, network string, request *QuoteReq) (*QuoteRes, int, error) {
	endpoint := "/quote"

	var queries = make(map[string]interface{})

	if request.FromTokenAddress == "" || request.ToTokenAddress == "" || request.Amount == "" {
		return nil, 0, fmt.Errorf("required parameter is missing")
	}

	_, _ = endpoint, queries

	return nil, 0, nil
}
