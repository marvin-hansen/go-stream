package api

import . "go-stream/api/types"

//Supported exchanges & assets
// https://www.coinapi.io/integration
// BINANCE			Binance
// BINANCEFTSCUAT	Binance Futures Testnet (Coin/fapi)
// BINANCEFTSC		Binance Futures (Coin/fapi)
// BINANCEFTS		Binance Futures (USDT/dapi)

func NewSDK(apiKey string) (sdk SDK) {
	validateApiKey(apiKey)
	config := getSDKConfig(apiKey)
	sdk = getSDK(config)
	return sdk
}
