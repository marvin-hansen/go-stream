package api

import . "go-stream/api/types"

const apiKey = "YOUR_API_KEY_HERE"

func NewSDK() (sdk SDK) {
	config := getSDKConfig(apiKey)
	sdk = getSDK(config)
	return sdk
}
