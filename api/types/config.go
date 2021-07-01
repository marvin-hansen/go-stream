package types

import "time"

type SdkConfig struct {
	ApiKey      string
	ApiVersion  ApiVersion
	Environment EnvironmentType
	Heartbeat   bool
}

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint           string
	WebsocketKeepalive bool
	WebsocketTimeout   time.Duration
}
