package v1

import (
	"github.com/gorilla/websocket"
	. "go-stream/api/types"
)

type SDKImpl struct {
	config *SdkConfig
}

var (
	con *websocket.Conn
)

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
	sdk = &SDKImpl{sdkConfig}
	sdk.init()
	return sdk
}

func (s SDKImpl) init() {
	wsConfig := s.getWSConfig()
	_ = s.connect(wsConfig)
}

func (s SDKImpl) SendHello(hello Hello) {

}
