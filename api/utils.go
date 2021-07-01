package api

import (
	. "go-stream/api/types"
	v1 "go-stream/api/v1"
)

func getSDK(sdkConfig *SdkConfig) (sdk SDK) {
	switch sdkConfig.ApiVersion {
	case ApiV1:
		// Bind interface to implementation with matching API version.
		sdk = v1.NewCoinApiSDKV1(sdkConfig)

	default:
		sdk = v1.NewCoinApiSDKV1(sdkConfig)
	}
	return sdk
}

func getSDKConfig(apiKey string) (sdkConfig *SdkConfig) {
	sdkConfig = &SdkConfig{
		ApiKey:      apiKey,
		ApiVersion:  ApiV1,
		Environment: TestInsecure,
		Heartbeat:   false,
	}
	return sdkConfig
}
