package go1inch

import "context"

// Get calldata for approve transaction and spender address
// Do not combine amount parameter with infinity parameter, only one must be sent.
// infinity will overwrite amount
// amount is set in minimal divisible units: for example, to unlock 1 DAI, amount should be 1000000000000000000, to unlock 1.03 USDC, amount should be 1030000.
func (c *Client) ApproveCalldata(ctx context.Context, network string, amount int64, infinity bool, tokenAddress string) (*ApproveCalldataResponse, int, error) {
	endpoint := "/approve/calldata"

	var queries = make(map[string]interface{})

	if infinity {
		queries["infinity"] = true
	} else {
		queries["amount"] = amount
	}
	queries["tokenAddress"] = tokenAddress

	var dataRes ApproveCalldataResponse
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
