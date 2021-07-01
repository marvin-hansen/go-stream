package v1

import (
	. "go-stream/api/types"
)

type SDKImpl struct {
	apiKey string
	url    string
}

func NewCoinApiSDKV1(apiKey string, envType EnvironmentType) (sdk *SDKImpl) {
	url := getUrl(envType)
	sdk = &SDKImpl{
		apiKey: apiKey,
		url:    url,
	}
	return sdk
}

func (s SDKImpl) SendHelloMessage(hello Hello) {

}

func (s SDKImpl) CloseConnect() {

}
