package go1inch

import "net/http"

type Client struct {
	Http *http.Client
}

type HealthcheckRes struct {
	Status string `json:"status"`
}

type ApproveCalldataReq struct {
	// Amount of tokens to be approved: (optional)
	// 0 — set approval to zero (lock a token)
	// >0 — approve exact amount of tokens.
	Amount int64
	// If set, approve infinite amount of tokens (optional)
	Infinity bool
	// The contract address of a token (required)
	TokenAddress string
}

type ApproveCalldataRes struct {
	// token contract address
	To []string `json:"to"`
	// amount of eth to be sent (in wei)
	Value string `json:"value"`
	// recommended gas price (in wei)
	GasPrice string `json:"gasPrice"`
	// result calldata
	Data string `json:"data"`
}

type ApproveSpenderRes struct {
	// address of 1inch contract
	Address string `json:"address"`
}

type QuoteReq struct {
	// contract address of a token to sell (required)
	FromTokenAddress string
	// contract address of a token to buy (required)
	ToTokenAddress string
	// amount of a token to sell (required)
	Amount string
	// referrer's fee in percentage (optional)
	Fee int64
	// liquidity protocols that can be used in a swap (optional)
	Protocols string
	// gas price (optional)
	GasPrice string
	// how many connectorTokens can be used (optional)
	ComplexityLevel string
	// contract addresses of connector tokens (optional)
	ConnectorTokens string
	// maximum amount of gas for a swap (optional)
	GasLimit int64
	// maximum number of parts each main route part can be split into (optional)
	Parts int64
	// maximum number of main route parts (optional)
	MainRouteParts int64
}

type Token struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"asddress"`
	Decimals int64  `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type Protocol struct {
	Name             string `json:"name"`
	Part             int64  `json:"part"`
	FromTokenAddress string `json:"fromTokenAddress"`
	ToTokenAddress   string `json:"toTokenAddress"`
}

type QuoteRes struct {
	// parameters of a token to sell
	FromToken Token `json:"fromToken"`
	// parameters of a token to buy
	ToToken Token `json:"ToToken"`
	// input amount of fromToken in minimal divisible units
	ToTokenAmount string `json:"toTokenAmount"`
	// result amount of toToken in minimal divisible units
	FromTokenAmount string `json:"fromTokenAmount"`
	// route of the trade
	Protocols []Protocol `json:"protocols"`
	// rough estimated amount of the gas limit for used protocols;
	// do not use estimatedGas from the quote method as the gas limit of a transaction
	EstimatedGas int64 `json:"estimatedGas"`
}

type SwapReq struct {
	// contract address of a token to sell (required)
	FromTokenAddress string
	// contract address of a token to buy (required)
	ToTokenAddress string
	// amount of a token to sell (required)
	Amount string
	// address of a seller (required)
	FromAddress string
	// additional slippage in percentage (required)
	Slippage string
	// protocols that can be used in a swap (optional)
	Protocols string
	// address that will receive a purchased token (optional)
	DestReceiver string
	// referrer's address (optional)
	ReferrerAddress string
	// referrer's fee in percentage (optional)
	Fee int64
	// gas price (optional)
	GasPrice string
	// if true, CHI will be burned from fromAddress to compensate gas (optional)
	BurnChi bool
	// how many connectorTokens can be used (optional)
	ComplexityLevel string
	// contract addresses of connector tokens (optional)
	ConnectorTokens string
	// if true, accept the partial order execution (optional)
	AllowPartialFill bool
	// if true, checks of the required quantities are disabled (optional)
	DisableEstimate bool
	// maximum amount of gas for a swap (optional)
	GasLimit int64
	// maximum number of parts each main route part can be split into (optional)
	Parts int64
	// maximum number of main route parts (optional)
	MainRouteParts int64
}

type Tx struct {
	// transactions will be sent from this address
	From string `json:"from"`
	// transactions will be sent to our contract address
	To string `json:"to"`
	// call data
	Data string `json:"data"`
	// amount of ETH (in wei) will be sent to the contract address
	Value string `json:"value"`
	// gas price in wei
	GasPrice string `json:"gasPrice"`
	// estimated amount of the gas limit, increase this value by 25%
	Gas int64 `json:"gas"`
}

type SwapRes struct {
	// parameters of a token to sell
	FromToken Token `json:"fromToken"`
	// parameters of a token to buy
	ToToken Token `json:"ToToken"`
	// input amount of fromToken in minimal divisible units
	ToTokenAmount string `json:"toTokenAmount"`
	// result amount of toToken in minimal divisible units
	FromTokenAmount string `json:"fromTokenAmount"`
	// route of the trade
	Protocols []Protocol `json:"protocols"`
	// transaction data
	Tx Tx `json:"tx"`
}
