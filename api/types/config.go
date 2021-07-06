package types

type SdkConfig struct {
	ApiKey      string
	ApiVersion  ApiVersion
	Environment EnvironmentType
}

// WsConfig webservice configuration
type WsConfig struct {
	ApiKey             string
	Endpoint           string
	WebsocketKeepalive bool
	WebsocketTimeout   int
}
