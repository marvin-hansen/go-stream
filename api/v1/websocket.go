package v1

import (
	"github.com/gorilla/websocket"
	"go-stream/api/types"
	"log"
	"time"
)

func (s SDKImpl) SendHello(hello types.Hello) {
	mtd := "SendHello: "

	b, err := hello.GetJSON()
	logError(err)

	err = con.WriteMessage(1, b)
	if err != nil {
		log.Println(mtd, "can't send Hello message!")
		logError(err)
		return
	}

	if running == false {
		_, stopC, err = s.process()
		if err != nil {
			log.Println("error processing messages")
			logError(err)
			running = false
			return
		}
		running = true
	}
}

func (s SDKImpl) process() (doneC, stopC chan struct{}, err error) {
	mtd := "process: "

	errHandler := logError
	handler := s.getWSHandler(errHandler)

	doneC = make(chan struct{})
	stopC = make(chan struct{})

	println(mtd, " * Processing ")

	go func() {
		// This function will exit either on error from ReadMessage
		// or when the stopC channel is closed by the client.
		defer close(doneC)

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

		var message []byte
		for {
			_, message, err = con.ReadMessage()
			if err != nil {
				if !silent {
					log.Println(mtd + "Error reading message!")
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
