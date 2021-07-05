package main

import (
	"fmt"
)

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
			  "apikey": "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C",
			  "heartbeat": false,
			  "subscribe_data_type": ["ohlcv"],
			  "subscribe_filter_asset_id": ["BTC", "ETH"]
			}
`

	return ""

}
