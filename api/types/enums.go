package types

type ApiVersion int

const (
	ApiV1 ApiVersion = iota
)

// EnvironmentType
// https://docs.coinapi.io/#endpoints-2
type EnvironmentType int

const (
	ProdEncrypted EnvironmentType = iota
	ProdInsecure
	TestEncrypted
	TestInsecure
)

// MessageType mimics a string enum.
type MessageType string

const (
	TRADE        MessageType = "trade"
	QUOTE        MessageType = "quote"
	BOOK         MessageType = "book"
	BOOK5        MessageType = "book5"
	BOOK20       MessageType = "book20"
	BOOK50       MessageType = "book50"
	OHLCV        MessageType = "ohlcv"
	VOLUME       MessageType = "volume"
	HEARTBEAT    MessageType = "heartbeat"
	HELLO        MessageType = "hello"
	ERROR        MessageType = "error"
	EXCHANGERATE MessageType = "exrate"
	RECONNECT    MessageType = "reconnect"
)
