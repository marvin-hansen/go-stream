package main

import (
	"encoding/json"
	"fmt"
	"go-stream/api"
	"time"
)

import SDK "github.com/CoinAPI/coinapi-sdk/data-api/go-rest/v1"

const apiKey = "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C"

func main() {
	//TestMetaListExchanges()
	//TestMetaListSymbols()

	TestHello()
}

func TestMetaListExchanges() {
	//https://github.com/coinapi/coinapi-sdk/blob/master/data-api/go-rest/main.go

	sdk := SDK.NewSDK(apiKey)

	exchanges, _ := sdk.Metadata_list_exchanges()
	fmt.Println("exchanges:")
	fmt.Println("  number:", len(exchanges))
	exchanges_item, _ := json.MarshalIndent(&exchanges[0], "", "  ")
	fmt.Println("  first items:", string(exchanges_item))

	for i, _ := range exchanges {
		exchange, _ := json.MarshalIndent(&exchanges[i], "", "  ")
		fmt.Println("  first items:", string(exchange))
	}
}

func TestMetaListSymbols() {
	sdk := SDK.NewSDK(apiKey)

	spots, _, _, _ := sdk.Metadata_list_symbols()
	fmt.Println("spots:")
	fmt.Println("number:", len(spots))

	//for i, _ := range spots {
	//	spots_item, _ := json.MarshalIndent(&spots[i], "", "  ")
	//	fmt.Println("first items:", string(spots_item))
	//}
}

func TestMetaListExchangeRate() {

	sdk := SDK.NewSDK(apiKey)
	exchange_rat_specific, _ := sdk.Exchange_rates_get_specific_rate("BTC", "USDT")
	fmt.Println("exchange_rat_specific:")
	exchange_rat_specific_item, _ := json.MarshalIndent(&exchange_rat_specific, "", "  ")
	fmt.Println("first items:", string(exchange_rat_specific_item))
}

func TestHello() {
	println("Start!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)

	println(" * SetErrorInvoke!")
	errInvokeFunc := GetErrorInvoke()
	sdk.SetErrorInvoke(errInvokeFunc)

	println(" * SetOHLCVInvoke!")
	OHLCVInvoke := GetOHLCVInvoke()
	sdk.SetOHLCVInvoke(OHLCVInvoke)

	println(" * SetTradesInvoke!")
	tradeInvoke := GetTradeInvoke()
	sdk.SetTradesInvoke(tradeInvoke)

	println(" * SetQuoteInvoke!")
	quoteInvoke := GetQuoteInvoke()
	sdk.SetQuoteInvoke(quoteInvoke)

	println(" * SetExRateInvoke!")
	exRateInvoke := GetExchangeInvoke()
	sdk.SetExRateInvoke(exRateInvoke)

	bookInvoke := GetBookInvoke()
	sdk.SetBookInvoke(bookInvoke)

	volInvoke := GetVolumeInvoke()
	sdk.SetVolumeInvoke(volInvoke)

	println(" * GetHello!")
	hello := getHello()

	println(" * SendHello!")
	sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 5)

	println(" * CloseConnection!")
	sdk.CloseConnection()

	println("Goodbye!")
}
