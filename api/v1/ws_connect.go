package v1

import (
	"github.com/gorilla/websocket"
	"go-stream/api/types"
	"log"
)

func (s SDKImpl) connect(wsConfig *types.WsConfig) (err error) {
	url := wsConfig.Endpoint

	con, _, err = websocket.DefaultDialer.Dial(wsConfig.Endpoint, nil)
	if err != nil {
		panic("Cannot connect to: " + url)
		return err
	}

	return nil
}

func (s SDKImpl) CloseConnection() {
	err := con.Close()
	if err != nil {
		log.Println("Can't close connection")
	}

}
