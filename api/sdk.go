package api

import (
	. "go-stream/api/types"
	v1 "go-stream/api/v1"
)

const (
	apiKey     = "YOUR_API_KEY_HERE"
	env        = TestInsecure
	apiVersion = ApiV1
)

func NewSDK() (sdk SDK) {
	// Bind interface to implementation with matching API version.
	sdk = getSDK(apiVersion)
	return sdk
}

func getSDK(version ApiVersion) (sdk SDK) {
	switch version {
	case ApiV1:
		return v1.NewCoinApiSDKV1(apiKey, env)

	default:
		return v1.NewCoinApiSDKV1(apiKey, env)
	}
}
