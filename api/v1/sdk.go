package v1

import (
	"github.com/gorilla/websocket"
	. "go-stream/api/types"
)

type SDKImpl struct {
	config *SdkConfig
}

var (
	con     *websocket.Conn
	stopC   chan struct{}
	running bool
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
)

func NewCoinApiSDKV1(sdkConfig *SdkConfig) (sdk *SDKImpl) {
	sdk = &SDKImpl{sdkConfig}
	sdk.init()
	return sdk
}

func (s SDKImpl) init() {
	wsConfig := s.getWSConfig()
	s.connect(wsConfig)
	running = false
}
