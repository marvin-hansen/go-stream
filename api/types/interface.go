package types

type SDK interface {
	SendHello(hello Hello)
	CloseConnection()
	SetErrorInvoke(function InvokeFunction)
	SetHeartBeatInvoke(function InvokeFunction)

	// Data handlers
	SetExRateInvoke(function InvokeFunction)
	SetTradesInvoke(function InvokeFunction)
	SetQuoteInvoke(function InvokeFunction)
	SetBookInvoke(function InvokeFunction)
	SetOHLCVInvoke(function InvokeFunction)
	SetVolumeInvoke(function InvokeFunction)
}
