package v1

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	t "go-stream/api/types"
)

func (s SDKImpl) processMessages() (doneC, stopC chan struct{}, err error) {

	errHandler := logError
	wsConfig := s.getWSConfig()
	wsHandler := s.getWSHandler(errHandler)

	return wsServe(wsConfig, wsHandler, errHandler)
}

func (s SDKImpl) getWSHandler(errHandler t.WsErrHandler) (wsHandler t.WsHandler) {
	wsHandler = func(message []byte) {
		err := s.processMessage(message, errHandler)
		if err != nil {
			errHandler(err)
		}
	}
	return wsHandler
}

func (s SDKImpl) processMessage(message []byte, errHandler t.WsErrHandler) (err error) {
	mtd := "processMessage: "
	println(mtd)

	var dataMessage = new(t.DataMessage)
	messageType := s.getMessageType(message, errHandler)

	switch messageType {
	case t.ERROR:
		// https://docs.coinapi.io/#error-handling
		errorMsg := new(t.ErrorMessage)
		_ = json.Unmarshal(message, errorMsg)
		dataMessage.ErrorMessage = errorMsg
		err = errorInvoke(dataMessage)
		return checkError(err)

	case t.EXCHANGERATE:
		// https://docs.coinapi.io/#exchange-rate-in
		msg := new(t.ExchangeRate)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage.ExchangeRate = msg
		err = exchangeInvoke(dataMessage)
		return checkError(err)

	case t.BOOK:
		// https://docs.coinapi.io/#orderbook-l2-full-in
		dataMessage = s.unMarshalOrderBook(message, messageType, errHandler)
		err = bookInvoke(dataMessage)
		return checkError(err)

	case t.BOOK5:
		dataMessage = s.unMarshalOrderBook(message, messageType, errHandler)
		err = bookInvoke(dataMessage)
		return checkError(err)

	case t.BOOK20:
		dataMessage = s.unMarshalOrderBook(message, messageType, errHandler)
		err = bookInvoke(dataMessage)
		return checkError(err)

	case t.BOOK50:
		dataMessage = s.unMarshalOrderBook(message, messageType, errHandler)
		err = bookInvoke(dataMessage)
		return checkError(err)

	case t.OHLCV:
		// https://docs.coinapi.io/#ohlcv-in
		msg := new(t.Ohlcv)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Ohlcv = msg
		err = ohlcvInvoke(dataMessage)
		return checkError(err)

	case t.QUOTE:
		// https://docs.coinapi.io/#quotes-in
		msg := new(t.Quote)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Quote = msg
		err = quotesInvoke(dataMessage)
		return checkError(err)

	case t.RECONNECT:
		// https://docs.coinapi.io/#reconnect-in
		msg := new(t.Reconnect)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Reconnect = msg
		err = reconnectInvoke(dataMessage)
		return checkError(err)

	case t.HEARTBEAT:
		// https://docs.coinapi.io/#heartbeat-in
		msg := new(t.Heartbeat)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Hearbeat = msg
		// call heartbeat here
		return nil

	case t.TRADE:
		// https://docs.coinapi.io/#trades-in
		msg := new(t.Trade)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Trade = msg
		err = tradesInvoke(dataMessage)
		return checkError(err)

	case t.VOLUME:
		// https://docs.coinapi.io/#volume-in
		msg := new(t.Volume)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Volume = msg
		err = volumeInvoke(dataMessage)
		return checkError(err)
	}
	return nil
}

func (s SDKImpl) unMarshalOrderBook(message []byte, msgType t.MessageType, errHandler t.WsErrHandler) (dataMessage *t.DataMessage) {
	bookMsg := new(t.OrderBook)
	bookMsg.Type = msgType
	err := json.Unmarshal(message, bookMsg)
	if err != nil {
		errHandler(err)
		return nil
	}

	dataMessage = new(t.DataMessage)
	dataMessage.Orderbook = bookMsg
	return dataMessage

}

func (s SDKImpl) getMessageType(message []byte, errHandler t.WsErrHandler) (messageType t.MessageType) {
	j, err := newJSON(message)
	if err != nil {
		errHandler(err)
		return
	}
	messageType = t.MessageType(j.Get("type").MustString())
	return messageType
}

const layout = "2009-01-11T15:04:05.0000000Z"

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}
