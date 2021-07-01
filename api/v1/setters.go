package v1

import (
	. "go-stream/api/types"
)

func (s SDKImpl) SetTradesInvoke(function InvokeFunction) {
	tradesInvoke = function
}

func (s SDKImpl) SetQuoteInvoke(function InvokeFunction) {
	quotesInvoke = function
}

func (s SDKImpl) SetBookInvoke(function InvokeFunction) {
	bookInvoke = function
}
func (s SDKImpl) SetOHLCVInvoke(function InvokeFunction) {
	ohlcvInvoke = function
}
func (s SDKImpl) SetVolumeInvoke(function InvokeFunction) {
	volumeInvoke = function
}
func (s SDKImpl) SetErrorInvoke(function InvokeFunction) {
	errorInvoke = function
}
func (s SDKImpl) SetReconnectInvoke(function InvokeFunction) {
	reconnectInvoke = function
}
