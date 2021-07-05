package main

import (
	"errors"
	"go-stream/api/types"
	"log"
)

func GetOHLCVInvoke() (errorInvoke types.InvokeFunction) {
	return func(message *types.DataMessage) (err error) {

		msg := message.Ohlcv
		log.Println(msg)

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
	var dataTypes = [25]string{}
	dataTypes[0] = "ohlcv"
	var assetIds = [25]string{}
	assetIds[0] = "BTC"

	hello = types.Hello{
		Type:                                    "hello",
		Api_key:                                 apiKey,
		Heartbeat:                               false,
		Subscribe_data_type:                     dataTypes,
		Subscribe_filter_symbol_id:              assetIds,
		Subscribe_update_limit_ms_quote:         0,
		Subscribe_update_limit_ms_book_snapshot: 0,
	}

	return hello

}
