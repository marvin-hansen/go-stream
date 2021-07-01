package v1

type SDK struct {
	apiKey string
	url    string
}

func NewSDK(apiKey string, envType EnvironmentType) (sdk *SDK) {
	url := getUrl(envType)
	sdk = &SDK{
		apiKey: apiKey,
		url:    url,
	}
	return sdk
}
