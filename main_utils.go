package main

import (
	"errors"
	"go-stream/api/types"
	"log"
)

func GetBookInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.Orderbook
		log.Println(msg)
		println()
		return nil
	}
}

func GetVolumeInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.Volume
		log.Println(msg)
		println()
		return nil
	}
}

func GetExchangeInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.ExchangeRate
		log.Println(msg)
		println()
		return nil
	}
}

func GetTradeInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.Trade
		log.Println(msg)
		println()
		return nil
	}
}

func GetQuoteInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.Quote
		log.Println(msg)
		println()
		return nil
	}
}
func GetOHLCVInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {
		msg := message.Ohlcv
		log.Println(msg)
		println()
		return nil
	}
}

func GetErrorInvoke() (errorInvoke types.InvokeFunction) {
	// You need to be prepared to receive an error message from us when you send something wrong;
	// all errors are permanent and you should expect that the underlying
	// WebSocket connection will be closed by us after sending an error message.
	// Good practice is to store all error messages somewhere for further manual review.
	// https://docs.coinapi.io/#error-handling
	return func(message *types.DataMessage) (err error) {
		mtd := "ErrorHandler: "
		println(mtd)
		msg := message.ErrorMessage.Message
		if msg != "" {
			log.Println(mtd+"ErrorMessage: ", msg)
			return errors.New(msg)
		}
		return nil
	}
}

func GetHeartBeatInvoke() (errorInvoke types.InvokeFunction) {
	// WebSocket working on TCP protocol which doesn’t have the feature indicating that the connection is broken without
	// trying exchange data over it. As you will not be actively sending messages to us,
	// with Heartbeat you can distinguish a situation where there are no market data updates
	// for you (Heartbeat is delivered but no market data updates)
	// or connection between us is broken (Heartbeat and market data updates are not delivered).
	// https://docs.coinapi.io/#reconnect-in
	return func(message *types.DataMessage) (err error) {
		log.Println()
		log.Println("===============================")
		log.Println("= Tracking heartbeat message! =")
		log.Println("===============================")
		log.Println()
		return nil
	}
}

func getHello(expanded, heartbeat bool) (hello types.Hello) {
	// After your WebSocket connection is established, you must send us a Hello message which contains:
	// * Stream preferences (Heartbeat and subscription details)
	// * API key for authorization
	// If your message will be incorrect, we will send you error message and disconnect connection afterwards.
	// Hello message can be repeated, each one will cause subscription scope override without interruption
	// of your WebSocket connection.
	// https://docs.coinapi.io/#hello-out

	var dataTypes []string
	var symbolIds []string
	var periodIDs []string

	dataTypes = append(dataTypes, "ohlcv")
	if expanded {
		dataTypes = append(dataTypes, "trade")
		dataTypes = append(dataTypes, "quote")
		dataTypes = append(dataTypes, "exrate")
		dataTypes = append(dataTypes, "book5")

	}

	symbolIds = append(symbolIds, "COINBASE_SPOT_BTC_USD")
	periodIDs = append(periodIDs, "1MIN")

	hello = types.Hello{
		Type:                       "hello",
		Api_key:                    apiKey,
		Heartbeat:                  heartbeat,
		Subscribe_data_type:        dataTypes,
		Subscribe_filter_period_id: periodIDs,
		Subscribe_filter_symbol_id: symbolIds,
	}
	return hello
}
