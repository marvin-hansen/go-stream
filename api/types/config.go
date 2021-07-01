package types

type SdkConfig struct {
	ApiKey      string
	ApiVersion  ApiVersion
	Environment EnvironmentType
	Heartbeat   bool
	Timeout     int
}

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint           string
	WebsocketKeepalive bool
	WebsocketTimeout   int
}
