package main

import (
	"go-stream/api"
	"time"
)

const apiKey = "550ECBDB-B1EF-42FE-8702-19CCAD9C2A7C"

func main() {
	//TestConnection()
	//
	TestHello()
}

func TestHello() {
	println("Start!")

	println(" * NewSDK!")
	sdk := api.NewSDK(apiKey)

	println(" * GetErrorInvoke!")
	errInvokeFunc := GetErrorInvoke()

	println(" * SetErrorInvoke!")
	sdk.SetErrorInvoke(errInvokeFunc)

	println(" * GetOHLCVInvoke!")
	OHLCVInvoke := GetOHLCVInvoke()

	println(" * SetOHLCVInvoke!")
	sdk.SetOHLCVInvoke(OHLCVInvoke)

	//println(" * GetHello!")
	//hello := getHello()

	//println(" * SendHello!")
	//sdk.SendHello(hello)

	println(" * Wait for messages!")
	time.Sleep(time.Second * 1)

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
