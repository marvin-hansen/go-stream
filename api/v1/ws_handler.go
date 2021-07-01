package v1

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	t "go-stream/api/types"
)

func (s SDKImpl) getWSHandler(errHandler t.WsErrHandler, invokeFunction t.InvokeFunction) (wsHandler t.WsHandler) {

	wsHandler = func(message []byte) {
		dataMsg := s.getDataMessage(message, errHandler)
		err := s.handle(dataMsg, invokeFunction)
		if err != nil {
			errHandler(err)
		}
	}
	return wsHandler
}

func (s SDKImpl) handle(message *t.DataMessage, invokeFunction t.InvokeFunction) (err error) {
	if invokeFunction == nil {
		return
	}
	return invokeFunction(message)
}

func (s SDKImpl) getDataMessage(message []byte, errHandler t.WsErrHandler) (dataMessage *t.DataMessage) {

	messageType := s.getMessageType(message, errHandler)

	switch messageType {
	case t.ERROR:
		// https://docs.coinapi.io/#error-handling
		errorMsg := new(t.ErrorMessage)
		_ = json.Unmarshal(message, errorMsg)
		dataMessage = new(t.DataMessage)
		dataMessage.ErrorMessage = errorMsg
		return dataMessage

	case t.EXCHANGERATE:
		// https://docs.coinapi.io/#exchange-rate-in
		msg := new(t.ExchangeRate)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.ExchangeRate = msg
		return dataMessage

	case t.BOOK:
		// https://docs.coinapi.io/#orderbook-l2-full-in
		return s.unMarshalOrderBook(message, messageType, errHandler)

	case t.BOOK5:
		return s.unMarshalOrderBook(message, messageType, errHandler)

	case t.BOOK20:
		return s.unMarshalOrderBook(message, messageType, errHandler)

	case t.BOOK50:
		return s.unMarshalOrderBook(message, messageType, errHandler)

	case t.OHLCV:
		// https://docs.coinapi.io/#ohlcv-in
		msg := new(t.Ohlcv)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Ohlcv = msg
		return dataMessage

	case t.QUOTE:
		// https://docs.coinapi.io/#quotes-in
		msg := new(t.Quote)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Quote = msg
		return dataMessage

	case t.RECONNECT:
		// https://docs.coinapi.io/#reconnect-in
		msg := new(t.Reconnect)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Reconnect = msg
		return dataMessage

	case t.HEARTBEAT:
		// https://docs.coinapi.io/#heartbeat-in
		msg := new(t.Heartbeat)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Hearbeat = msg
		return dataMessage

	case t.TRADE:
		// https://docs.coinapi.io/#trades-in
		msg := new(t.Trade)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Trade = msg
		return dataMessage

	case t.VOLUME:
		// https://docs.coinapi.io/#volume-in
		msg := new(t.Volume)
		msg.Type = messageType
		_ = json.Unmarshal(message, msg)
		dataMessage = new(t.DataMessage)
		dataMessage.Volume = msg
		return dataMessage
	}

	return dataMessage
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
