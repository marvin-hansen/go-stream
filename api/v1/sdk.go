package v1

import (
	. "go-stream/api/types"
	"go-stream/api/web_socket"
)

type SDKImpl struct {
	config *SdkConfig
	ws     *web_socket.WebSocket
}

var (
	running  bool
	helloMsg *Hello
)

var (
	// data handlers
	tradesInvoke   InvokeFunction
	quotesInvoke   InvokeFunction
	bookInvoke     InvokeFunction
	ohlcvInvoke    InvokeFunction
	volumeInvoke   InvokeFunction
	exchangeInvoke InvokeFunction
	// sys handlers
	errorInvoke     InvokeFunction
	heartBeatInvoke InvokeFunction
	reconnectInvoke InvokeFunction
)

func NewCoinApiSDKV1(sdkConfig *SdkConfig) (sdk *SDKImpl) {
	sdk = new(SDKImpl)
	sdk.init(sdkConfig)
	return sdk
}

func (s SDKImpl) init(sdkConfig *SdkConfig) {
	s.config = sdkConfig
	url := getUrl(s.config.EnvironmentType)
	ws := web_socket.NewWebSocket(url)
	s.ws = ws
	//_ = s.OpenConnection() // errors get handled inside connect()
}
