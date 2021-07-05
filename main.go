package main

import (
	"go-stream/api"
	"time"
)

const apiKey = "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C"

func main() {
	TestConnection()
	//
	//TestHello()
}

func TestHello() {
	println("Start!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)

	println(" * SetErrorInvoke!")
	sdk.SetErrorInvoke(nil)

	println(" * SetOHLCVInvoke!")
	sdk.SetOHLCVInvoke(nil)

	println(" * SendHello!")
	hello := getHello()
	sdk.SendHello(hello)

	//time.Sleep(time.Second * 5)

	println(" * CloseConnection!")
	sdk.CloseConnection()

	println("Goodbye!")
}

func TestConnection() {
	println("Start!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)

	println(" * Connect & Wait!")
	time.Sleep(time.Second * 2)

	println(" * CloseConnection!")
	sdk.CloseConnection()
}
