package api

import (
	. "go-stream/api/types"
	v1 "go-stream/api/v1"
	"strings"
)

func validateApiKey(apiKey string) {
	if apiKey == "" {
		printHeaader("API key is empty!", "Add a valid API key to main!")
		panic("Invalid API KEY - Abort!")
	}

	if len(apiKey) <= 21 {
		printHeaader("API key is too short. Should be 30 or more characters long!", "Add a valid API key to main!")
		panic("Invalid API KEY - Abort!")
	}

	if strings.Contains(apiKey, "SAMPLE-KEY") {
		printHeaader("Inactive sample key detected!", "Add your own valid API key to main!")
		panic("Invalid API KEY - Abort!")
	}
}

func printHeaader(problem, solution string) {
	println()
	println("Problem: ", problem)
	println("Solution: ", solution)
	println()
}

func getSDKConfig(apiKey string) (sdkConfig *SdkConfig) {
	sdkConfig = &SdkConfig{
		ApiKey:      apiKey,
		ApiVersion:  ApiV1,
		Environment: TestInsecure,
	}
	return sdkConfig
}

func getSDK(sdkConfig *SdkConfig) (sdk SDK) {
	switch sdkConfig.ApiVersion {
	case ApiV1:
		// Bind interface to implementation matching the selected API version.
		sdk = v1.NewCoinApiSDKV1(sdkConfig)

	default:
		sdk = v1.NewCoinApiSDKV1(sdkConfig)
	}
	return sdk
}
