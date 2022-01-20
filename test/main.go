package main

import (
	"fmt"
)

const API_KEY = ""

func main() {
	println("Hello Main!")

	hello := getHello()

	println(hello)

	fmt.Println("Goodbye!")
}

func getHello() (hello string) {

	hello = `
			{
			  "type": "hello",
			  "apikey": API_KEY,
			  "heartbeat": false,
			  "subscribe_data_type": ["ohlcv"],
			  "subscribe_filter_asset_id": ["BTC", "ETH"]
			}
`

	return ""

}
