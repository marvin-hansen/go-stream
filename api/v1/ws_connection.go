package v1

import (
	"github.com/gorilla/websocket"
	"go-stream/api/types"
	"log"
)

func (s SDKImpl) connect(wsConfig *types.WsConfig) {
	mtd := "connect: "
	url := wsConfig.Endpoint
	var err error
	con, _, err = websocket.DefaultDialer.Dial(wsConfig.Endpoint, nil)
	if err != nil {
		log.Println(mtd, err)
		panic(mtd + "Cannot connect to: " + url)
	}
}

func (s SDKImpl) CloseConnection() {
	mtd := "CloseConnection: "

	// Stop processing messages
	running = false

	// close WS channel if its not yet fully closed!
	if stopC != nil {
		close(stopC)
	}

	// close connection
	err := con.Close()
	if err != nil {
		log.Println(mtd + "Can't close connection")
		log.Println(err)
	}
}
