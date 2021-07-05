package api

import (
	. "go-stream/api/types"
	v1 "go-stream/api/v1"
)

func getSDK(sdkConfig *SdkConfig) (sdk SDK) {
	switch sdkConfig.ApiVersion {
	case ApiV1:
		// Bind interface to implementation matching API version.
		sdk = v1.NewCoinApiSDKV1(sdkConfig)

	default:
		sdk = v1.NewCoinApiSDKV1(sdkConfig)
	}
	return sdk
}

func getSDKConfig(apiKey string) (sdkConfig *SdkConfig) {

	var exchanges = make([]string, 3)
	exchanges[0] = "BINANCE"

	sdkConfig = &SdkConfig{
		ApiKey:      apiKey,
		ApiVersion:  ApiV1,
		Environment: TestInsecure,
		Exchanges:   exchanges,
		Heartbeat:   false,
		Timeout:     3,
	}
	return sdkConfig
}
