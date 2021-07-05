package types

type SDK interface {
	SendHello(hello Hello)
	CloseConnection()
	SetErrorInvoke(function InvokeFunction)
	SetExRateInvoke(function InvokeFunction)
	SetTradesInvoke(function InvokeFunction)
	SetQuoteInvoke(function InvokeFunction)
	SetBookInvoke(function InvokeFunction)
	SetOHLCVInvoke(function InvokeFunction)
	SetVolumeInvoke(function InvokeFunction)
}
