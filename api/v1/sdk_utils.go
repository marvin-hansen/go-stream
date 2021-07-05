package v1

import "go-stream/api/types"

func (s SDKImpl) getWSConfig() (wsCfg *types.WsConfig) {
	url := getUrl(s.config.Environment)
	wsCfg = &types.WsConfig{
		ApiKey:             s.config.ApiKey,
		Endpoint:           url,
		WebsocketKeepalive: s.config.Heartbeat,
		WebsocketTimeout:   s.config.Timeout,
	}
	return wsCfg
}
