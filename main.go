package main

import (
	"go-stream/api"
	"go-stream/api/types"
	"time"
)

const apiKey = "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C"

func main() {
	TestSendHello()
	TestHeartbeat()
	TestReconnect()
}

func TestSendHello() {

	printHeader("TestSendHello!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)
	// verbose switches off console print of heartbeat & reconnect messages
	sys := getSysInvokes(false)

	println(" * SetErrorInvoke!")
	sdk.SetErrorInvoke(sys.ErrorInvoke)

	println(" * SetHeartBeatInvoke!")
	sdk.SetHeartBeatInvoke(sys.HeartBeatInvoke)

	println(" * SetReconnectInvoke!")
	sdk.SetReconnectInvoke(sys.ReconnectInvoke)

	println(" * SetOHLCVInvoke!")
	OHLCVInvoke := GetInvokeFunction(types.OHLCV)
	sdk.SetOHLCVInvoke(OHLCVInvoke)

	println(" * SetTradesInvoke!")
	tradeInvoke := GetInvokeFunction(types.TRADE)
	sdk.SetTradesInvoke(tradeInvoke)

	println(" * SetQuoteInvoke!")
	quoteInvoke := GetInvokeFunction(types.QUOTE)
	sdk.SetQuoteInvoke(quoteInvoke)

	println(" * SetExRateInvoke!")
	exRateInvoke := GetInvokeFunction(types.EXCHANGERATE)
	sdk.SetExRateInvoke(exRateInvoke)

	println(" * SetBookInvoke!")
	bookInvoke := GetInvokeFunction(types.BOOK_L2_FULL)
	sdk.SetBookInvoke(bookInvoke)

	// Volume throws a strange error
	println(" * GetVolumeInvoke!")
	volInvoke := GetInvokeFunction(types.VOLUME)
	sdk.SetVolumeInvoke(volInvoke)

	println(" * GetHello: Single data type!")
	hello := getHello(false, false)

	println(" * SendHello: Single data type!")
	_ = sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 5)

	println(" * GetHello: Expanded data types!")
	hello = getHello(true, false)

	println(" * SendHello: Expanded data types!")
	_ = sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 5)

	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")
}

func TestReconnect() {
	// This test only process OHLCV as no other invoke functions are set!
	printHeader("TestReconnect!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)
	sys := getSysInvokes(true)

	println(" * SetErrorInvoke!")
	sdk.SetErrorInvoke(sys.ErrorInvoke)

	println(" * SetHeartBeatInvoke!")
	sdk.SetHeartBeatInvoke(sys.HeartBeatInvoke)

	println(" * SetReconnectInvoke!")
	sdk.SetReconnectInvoke(sys.ReconnectInvoke)

	println(" * SetOHLCVInvoke!")
	OHLCVInvoke := GetInvokeFunction(types.OHLCV)
	sdk.SetOHLCVInvoke(OHLCVInvoke)

	println(" * GetHello: Single data type!")
	hello := getHello(false, true)

	println(" * SendHello: Single data type!")
	_ = sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 10)

	println("******************")
	println("* Hard Reconnect *")
	println("******************")
	_ = sdk.Reconnect()

	println(" * Wait for messages!")
	time.Sleep(time.Second * 5)

	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")
}

func TestHeartbeat() {
	// This test only process OHLCV as no other invoke functions are set!
	printHeader("TestHeartbeat!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)
	sys := getSysInvokes(true)

	println(" * SetErrorInvoke!")
	sdk.SetErrorInvoke(sys.ErrorInvoke)

	println(" * SetHeartBeatInvoke!")
	sdk.SetHeartBeatInvoke(sys.HeartBeatInvoke)

	println(" * SetReconnectInvoke!")
	sdk.SetReconnectInvoke(sys.ReconnectInvoke)

	println(" * SetOHLCVInvoke!")
	OHLCVInvoke := GetInvokeFunction(types.OHLCV)
	sdk.SetOHLCVInvoke(OHLCVInvoke)

	println(" * GetHello: Heartbeat!")
	hello := getHello(false, true)

	println(" * SendHello: Heartbeat!")
	_ = sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 10)

	println(" * CloseConnection!")
	_ = sdk.CloseConnection()

	println("Goodbye!")
}

func printHeader(msg string) {
	println()
	println("=====================")
	println("Start: " + msg)
	println("=====================")
	println()
}
