package v1

import (
	"github.com/gorilla/websocket"
	. "go-stream/api/types"
	"log"
	"time"
)

type SDKImpl struct {
	config *SdkConfig
}

const (
	// Time allowed to write to the web socket.
	writeWait = 1 * time.Second
)

var (
	con   *websocket.Conn
	stopC chan struct{}
)

var (
	tradesInvoke    InvokeFunction
	quotesInvoke    InvokeFunction
	bookInvoke      InvokeFunction
	ohlcvInvoke     InvokeFunction
	volumeInvoke    InvokeFunction
	errorInvoke     InvokeFunction
	exchangeInvoke  InvokeFunction
	reconnectInvoke InvokeFunction
)

func NewCoinApiSDKV1(sdkConfig *SdkConfig) (sdk *SDKImpl) {
	sdk = &SDKImpl{sdkConfig}
	sdk.init()
	return sdk
}

func (s SDKImpl) init() {

	// fetch config
	wsConfig := s.getWSConfig()

	// connect
	err := s.connect(wsConfig)
	log.Println(err)

	// process
	_, stopC, err = s.processMessages()
	if err != nil {
		return
	}
}
