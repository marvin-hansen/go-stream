package types

type SDK interface {
	SendHelloMessage(hello Hello)
	CloseConnect()
	SetTradesInvoke(function InvokeFunction)
	SetQuoteInvoke(function InvokeFunction)
	SetBookInvoke(function InvokeFunction)
	SetOHLCVInvoke(function InvokeFunction)
	SetVolumeInvoke(function InvokeFunction)
	SetErrorInvoke(function InvokeFunction)
	SetReconnectInvoke(function InvokeFunction)
}
