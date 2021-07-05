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

func (s SDKImpl) getHello(dataTypes, assetIds []string) (hello types.Hello) {
	//var dataTypes = make([]string, 10)
	//dataTypes[0] = "ohlcv"
	//var assetIds = make([]string, 10)
	//assetIds[0] = "BTC"

	hello = types.Hello{
		Type:                                    "hello",
		Api_key:                                 s.config.ApiKey,
		Heartbeat:                               s.config.Heartbeat,
		Subscribe_data_type:                     dataTypes,
		Subscribe_filter_symbol_id:              nil,
		Subscribe_filter_asset_id:               assetIds,
		Subscribe_filter_period_id:              nil,
		Subscribe_filter_exchange_id:            nil,
		Subscribe_update_limit_ms_quote:         0,
		Subscribe_update_limit_ms_book_snapshot: 0,
	}

	return hello

}