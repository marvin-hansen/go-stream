package v1

func getUrl(env EnvironmentType) (url string) {
	switch env {
	case ProdEncrypted:
		return ProductionEncrypted
	case ProdInsecure:
		return ProductionInsecure
	case TestEncrypted:
		return SandboxEncrypted
	case TestInsecure:
		return SandboxInsecure
	default:
		return SandboxEncrypted
	}

}

// EnvironmentType
// https://docs.coinapi.io/#endpoints-2
type EnvironmentType int

const (
	ProdEncrypted EnvironmentType = iota
	ProdInsecure
	TestEncrypted
	TestInsecure
)

// Endpoints WebSocket endpoint provides real-time market data streaming which works in Subscribe-Publish communication model.
// https://docs.coinapi.io/#md-websocket-api
const (
	ProductionEncrypted = "wss://ws.coinapi.io/v1/"
	ProductionInsecure  = "ws://ws.coinapi.io/v1/"
	SandboxEncrypted    = "wss://ws-sandbox.coinapi.io/v1/"
	SandboxInsecure     = "ws://ws-sandbox.coinapi.io/v1/"
)
