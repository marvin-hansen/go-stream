package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// :8000
var addr = flag.String("addr", "", "http service address")
var channel = ""

func main() {
	println("Main")

	flag.Parse()

	client, err := NewWebSocketClient(*addr, channel)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting")

	helloMsg := getHello()

	err = client.Write(helloMsg)
	if err != nil {
		fmt.Printf("error: %v, writing error\n", err)
	}

	//client.listen()

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
	fmt.Println("Goodbye")
}

func getHello() (hello string) {

	hello = `
			{
			  "type": "hello",
			  "apikey": "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C",
			  "heartbeat": false,
			  "subscribe_data_type": ["ohlcv"],
			  "subscribe_filter_asset_id": ["BTC", "ETH"]
			}
`

	return ""

}
