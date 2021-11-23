package go1inch

import (
	"context"
	"errors"
)

// ApproveTransaction gets calldata for approve transaction and spender address
// Do not combine amount parameter with infinity parameter, only one must be sent.
// infinity will overwrite amount
// amount is set in minimal divisible units: for example, to unlock 1 DAI, amount should be 1000000000000000000, to unlock 1.03 USDC, amount should be 1030000.
func (c *Client) ApproveTransaction(ctx context.Context, network, tokenAddress string, opts *ApproveTransactionOpts) (*ApproveTransactionRes, int, error) {
	endpoint := "/approve/transaction"
	if tokenAddress == "" {
		return nil, 0, errors.New("required parameter is missing")
	}

	var queries = make(map[string]interface{})

	queries["tokenAddress"] = tokenAddress

	if opts != nil {
		if opts.Amount != "" {
			queries["amount"] = opts.Amount
		}
	}

	var dataRes ApproveTransactionRes
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

// ApproveAllowance gets the number of tokens that the 1inch router is allowed to spend
func (c *Client) ApproveAllowance(ctx context.Context, network, tokenAddress, walletAddress string) (*ApproveAllowanceRes, int, error) {
	endpoint := "/approve/allowance"
	if tokenAddress == "" || walletAddress == "" {
		return nil, 0, errors.New("required parameter is missing")
	}

	var queries = make(map[string]interface{})

	queries["tokenAddress"] = tokenAddress
	queries["walletAddress"] = walletAddress
	var dataRes ApproveAllowanceRes
	statusCode, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
