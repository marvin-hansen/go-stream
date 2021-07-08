package v1

import (
	"log"
)

func (s SDKImpl) OpenConnection() (err error) {
	url := getUrl(s.config.EnvironmentType)
	err = s.ws.Connect(url, nil)
	return err
}

func (s SDKImpl) CloseConnection() (err error) {
	// Stop processing messages
	running = false

	// close connection
	err = s.ws.Close()
	if err != nil {
		log.Println("Can't close connection")
		log.Println(err)
	}
	return err
}

func (s SDKImpl) Reconnect() (err error) {

	err = s.CloseConnection()
	logError(err)

	err = s.OpenConnection()
	logError(err)

	hello := s.getHelloMessage()

	err = s.SendHello(hello)
	logError(err)

	return err
}
