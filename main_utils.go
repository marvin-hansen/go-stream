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

func getHello() (hello types.Hello) {
	var dataTypes []string
	var symbolIds []string
	var periodIDs []string

	dataTypes = append(dataTypes, "ohlcv")
	dataTypes = append(dataTypes, "trade")
	//dataTypes = append(dataTypes, "quote")
	//dataTypes = append(dataTypes, "exrate")
	//dataTypes = append(dataTypes, "book5")

	symbolIds = append(symbolIds, "COINBASE_SPOT_BTC_USD")
	periodIDs = append(periodIDs, "1MIN")

	hello = types.Hello{
		Type:                       "hello",
		Api_key:                    apiKey,
		Heartbeat:                  false,
		Subscribe_data_type:        dataTypes,
		Subscribe_filter_period_id: periodIDs,
		Subscribe_filter_symbol_id: symbolIds,
	}
	return hello
}
