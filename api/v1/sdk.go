package v1

import (
	. "go-stream/api/types"
)

type SDKImpl struct {
	apiKey    string
	url       string
	keepAlive bool
}

var (
	tradesInvoke    InvokeFunction
	quotesInvoke    InvokeFunction
	bookInvoke      InvokeFunction
	ohlcvInvoke     InvokeFunction
	volumeInvoke    InvokeFunction
	errorInvoke     InvokeFunction
	reconnectInvoke InvokeFunction
)

func NewCoinApiSDKV1(sdkConfig *SdkConfig) (sdk *SDKImpl) {
	url := getUrl(sdkConfig.Environment)
	sdk = &SDKImpl{
		apiKey:    sdkConfig.ApiKey,
		url:       url,
		keepAlive: sdkConfig.Heartbeat,
	}
	return sdk
}

func (s SDKImpl) SendHelloMessage(hello Hello) {

}

func (s SDKImpl) CloseConnect() {

}
