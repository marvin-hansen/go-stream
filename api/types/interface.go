package types

type SDK interface {
	SendHello(hello *Hello) (err error)
	OpenConnection() (err error)
	CloseConnection() (err error)
	// Reconnect closes the current connection, opens a new one,
	// and resends the last hello message. No message buffering, just hard reconnect!
	Reconnect() (err error)

	// sys handlers
	SetErrorInvoke(function InvokeFunction)
	SetHeartBeatInvoke(function InvokeFunction)
	SetReconnectInvoke(function InvokeFunction)

	// Data handlers
	SetExRateInvoke(function InvokeFunction)
	SetTradesInvoke(function InvokeFunction)
	SetQuoteInvoke(function InvokeFunction)
	SetBookInvoke(function InvokeFunction)
	SetOHLCVInvoke(function InvokeFunction)
	SetVolumeInvoke(function InvokeFunction)
}
