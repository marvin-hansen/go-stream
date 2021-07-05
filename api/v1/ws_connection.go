package v1

import (
	"github.com/gorilla/websocket"
	"go-stream/api/types"
	"log"
)

func (s SDKImpl) connect(wsConfig *types.WsConfig) {
	url := wsConfig.Endpoint
	var err error
	con, _, err = websocket.DefaultDialer.Dial(wsConfig.Endpoint, nil)
	if err != nil {
		panic("Cannot connect to: " + url)
	}
}

func (s SDKImpl) CloseConnection() {

	// close WS channel if its not yet fully closed!
	if stopC != nil {
		close(stopC)
	}

	// close connection
	err := con.Close()
	if err != nil {
		log.Println("Can't close connection")
	}
}
