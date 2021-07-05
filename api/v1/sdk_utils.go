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

func (s SDKImpl) getHello(dataTypes, symbolIDs, assetIDs, periodIDs []string) (hello types.Hello) {

	hello = types.Hello{
		Type:                         "hello",
		Api_key:                      s.config.ApiKey,
		Heartbeat:                    s.config.Heartbeat,
		Subscribe_data_type:          dataTypes,
		Subscribe_filter_symbol_id:   symbolIDs,
		Subscribe_filter_asset_id:    assetIDs,
		Subscribe_filter_period_id:   periodIDs,
		Subscribe_filter_exchange_id: s.config.Exchanges,
	}

	return hello

}
