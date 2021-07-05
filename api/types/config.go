package types

type SdkConfig struct {
	ApiKey      string
	ApiVersion  ApiVersion
	Environment EnvironmentType
	Exchanges   []string
	Heartbeat   bool
	Timeout     int
}

// WsConfig webservice configuration
type WsConfig struct {
	ApiKey             string
	Endpoint           string
	WebsocketKeepalive bool
	WebsocketTimeout   int
}