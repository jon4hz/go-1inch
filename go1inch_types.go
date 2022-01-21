package go1inch

import "net/http"

type Client struct {
	Http *http.Client
}

type HealthcheckRes struct {
	Status string `json:"status"`
}

type ApproveTransactionOpts struct {
	// Amount of tokens to be approved:
	// 0 — set approval to zero (lock a token)
	// >0 — approve exact amount of tokens.
	// if not set use infinity approve
	Amount string
}

type ApproveTransactionRes struct {
	// token contract address
	To string `json:"to"`
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

type ApproveAllowanceRes struct {
	// address of 1inch contract
	Allowance string `json:"allowance"`
}

type QuoteOpts struct {
	// referrer's fee in percentage
	// Ethereum: min: 0; max: 3; Binance: min: 0; max: 3; default: 0; !should be the same for quote and swap!
	Fee string
	// liquidity protocols that can be used in a swap
	Protocols string
	// gas price
	// default: fast from network
	GasPrice string
	// how many connectorTokens can be used
	// min: 0; max: 3; default: 2; !should be the same for quote and swap!
	ComplexityLevel string
	// contract addresses of connector tokens
	// max: 5; !should be the same for quote and swap!
	ConnectorTokens string
	// maximum amount of gas for a swap
	GasLimit string
	// virtual split parts. default: 50; max: 500; !should be the same for quote and swap!
	VirtualParts string
	// maximum number of parts each main route part can be split into
	// split parts. default: Ethereum: 50; Binance: 40 max: Ethereum: 100; Binance: 100; !should be the same for quote and swap!
	Parts string
	// maximum number of main route parts
	// default: Ethereum: 10, Binance: 10; max: Ethereum: 50, Binance: 50 !should be the same for quote and swap!
	MainRouteParts string
}

type Token struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Decimals int64  `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type Protocol [][]struct {
	Name             string  `json:"name"`
	Part             float64 `json:"part"`
	FromTokenAddress string  `json:"fromTokenAddress"`
	ToTokenAddress   string  `json:"toTokenAddress"`
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

type SwapOpts struct {
	// protocols that can be used in a swap
	Protocols string
	// address that will receive a purchased token
	// Receiver of destination currency. default: fromAddress
	DestReceiver string
	// referrer's address
	ReferrerAddress string
	// referrer's fee in percentage
	// Ethereum: min: 0; max: 3; Binance: min: 0; max: 3; default: 0; !should be the same for quote and swap!
	Fee string
	// gas price
	// default: fast from network
	GasPrice string
	// if true, CHI will be burned from fromAddress to compensate gas
	// default: false; Suggest to check user's balance and allowance before set this flag; CHI should be approved to spender address
	BurnChi bool
	// how many connectorTokens can be used
	// min: 0; max: 3; default: 2; !should be the same for quote and swap!
	ComplexityLevel string
	// contract addresses of connector tokens
	// max: 5; !should be the same for quote and swap!
	ConnectorTokens string
	// if true, accept the partial order execution
	AllowPartialFill bool
	// if true, checks of the required quantities are disabled
	DisableEstimate bool
	// maximum amount of gas for a swap
	GasLimit string
	// virtual split parts. default: 50; max: 500; !should be the same for quote and swap!
	VirtualParts string
	// maximum number of parts each main route part can be split into
	// split parts. default: Ethereum: 50; Binance: 40 max: Ethereum: 100; Binance: 100; !should be the same for quote and swap!
	Parts string
	// maximum number of main route parts
	// default: Ethereum: 10, Binance: 10; max: Ethereum: 50, Binance: 50 !should be the same for quote and swap!
	MainRouteParts string
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

type Protocols struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Img   string `json:"img"`
}

type LiquiditySoucesRes struct {
	Protocols []Protocols `json:"protocols"`
}

type TokensRes struct {
	Tokens map[string]Token `json:"tokens"`
}
