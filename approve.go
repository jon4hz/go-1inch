package go1inch

import "context"

// ApproveCalldata gets calldata for approve transaction and spender address
// Do not combine amount parameter with infinity parameter, only one must be sent.
// infinity will overwrite amount
// amount is set in minimal divisible units: for example, to unlock 1 DAI, amount should be 1000000000000000000, to unlock 1.03 USDC, amount should be 1030000.
func (c *Client) ApproveCalldata(ctx context.Context, network string, request *ApproveCalldataReq) (*ApproveCalldataRes, int, error) {
	endpoint := "/approve/calldata"

	var queries = make(map[string]interface{})

	if request.Infinity {
		queries["infinity"] = true
	} else {
		queries["amount"] = request.Amount
	}
	queries["tokenAddress"] = request.TokenAddress

	var dataRes ApproveCalldataRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// ApproveSpender gets the address to which you need to approve before the swap transaction
func (c *Client) ApproveSpender(ctx context.Context, network string) (*ApproveSpenderRes, int, error) {
	endpoint := "/approve/spender"
	var dataRes ApproveSpenderRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
