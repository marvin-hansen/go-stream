package v1

import (
	"go-stream/api/types"
	"log"
)

func (s SDKImpl) SendHello(hello *types.Hello) (err error) {
	b, err := hello.GetJSON()
	logError(err)

	err = s.ws.WriteByteMessage(b)
	if err != nil {
		log.Println("can't send Hello message!")
		logError(err)
		return err
	}

	s.setHelloMessage(hello) // store last hello message to restore after a hard re-connect.
	if running == false {
		err = s.startProcessing()
		if err != nil {
			log.Println("can't start message processing!")
			logError(err)
			return err
		}
		running = true
	}
	return err
}

func (s SDKImpl) startProcessing() (err error) {
	errHandler := logError
	handler := s.getWSHandler(errHandler)
	err = s.ws.ReadByteMessages(handler, errHandler)
	if err != nil {
		log.Println("error starting message processing!")
		logError(err)
		running = false
		return err
	}
	running = true
	return nil
}
