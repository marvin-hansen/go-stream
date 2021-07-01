package v1

import (
	. "go-stream/api/types"
)

type SDKImpl struct {
	apiKey    string
	url       string
	keepAlive bool
	timeOut   int
}

var (
	tradesInvoke    InvokeFunction
	quotesInvoke    InvokeFunction
	bookInvoke      InvokeFunction
	ohlcvInvoke     InvokeFunction
	volumeInvoke    InvokeFunction
	errorInvoke     InvokeFunction
	reconnectInvoke InvokeFunction
	wsCfg           WsConfig
)

func NewCoinApiSDKV1(sdkConfig *SdkConfig) (sdk *SDKImpl) {
	url := getUrl(sdkConfig.Environment)
	sdk = &SDKImpl{
		apiKey:    sdkConfig.ApiKey,
		url:       url,
		keepAlive: sdkConfig.Heartbeat,
		timeOut:   sdkConfig.Timeout,
	}
	sdk.init()
	return sdk
}

func (s SDKImpl) init() {
	wsCfg = WsConfig{
		Endpoint:           s.url,
		WebsocketKeepalive: s.keepAlive,
		WebsocketTimeout:   s.timeOut,
	}
}

func (s SDKImpl) SendHelloMessage(hello Hello) {

}

func (s SDKImpl) CloseConnect() {

}
