package v1

import (
	"go-stream/api/types"
	"log"
)

func (s SDKImpl) SendHello(hello types.Hello) {

	b, err := hello.GetJSON()
	logError(err)

	err = con.WriteMessage(1, b)
	if err != nil {
		log.Println("can't send Hello message!")
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

	errHandler := logError
	handler := s.getWSHandler(errHandler)

	doneC = make(chan struct{})
	stopC = make(chan struct{})

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
					log.Println("Error reading message!")
					errHandler(err)
				}
				return
			}

			//printRawMsg(message)
			handler(message)
		}
	}()
	return
}

func printRawMsg(message []byte) {
	msg := string(message)
	log.Println("message: ")
	log.Println(msg)
}
