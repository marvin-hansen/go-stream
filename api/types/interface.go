package types

type SDK interface {
	SendHello(hello Hello)
	CloseConnection()
	SetTradesInvoke(function InvokeFunction)
	SetQuoteInvoke(function InvokeFunction)
	SetBookInvoke(function InvokeFunction)
	SetOHLCVInvoke(function InvokeFunction)
	SetVolumeInvoke(function InvokeFunction)
	SetErrorInvoke(function InvokeFunction)
	SetReconnectInvoke(function InvokeFunction)
}
