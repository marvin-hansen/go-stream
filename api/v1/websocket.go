package v1

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-stream/api/types"
	"time"
)

func (s SDKImpl) SendHello(hello types.Hello) {

	b, err := json.Marshal(hello)
	logError(err)

	err = con.WriteMessage(1, b)
	logError(err)
}

var wsServe = func(cfg *types.WsConfig, handler types.WsHandler, errHandler types.WsErrHandler) (doneC, stopC chan struct{}, err error) {
	//con, _, err = websocket.DefaultDialer.Dial(cfg.Endpoint, nil)
	//if err != nil {
	//	return nil, nil, err
	//}

	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		//
		//if cfg.WebsocketKeepalive {
		//	keepAlive(con, time.Duration(cfg.WebsocketTimeout))
		//}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking  operation.

		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			_ = con.Close()
		}()
		for {
			_, message, err := con.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(5 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				_ = c.Close()
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				_ = c.Close()
				return
			}
		}
	}()
}

func closeConnection(c *websocket.Conn) (err error) {
	err = c.Close()
	if err != nil {
		return err
	} else {
		return nil
	}
}
