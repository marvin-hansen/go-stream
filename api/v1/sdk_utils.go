package v1

import "go-stream/api/types"

func (s SDKImpl) getWSConfig() (wsCfg *types.WsConfig) {
	url := getUrl(s.config.Environment)
	wsCfg = &types.WsConfig{
		ApiKey:   s.config.ApiKey,
		Endpoint: url,
	}
	return wsCfg
}
