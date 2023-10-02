package config

const (
	EthChainId   = 1
	BSCChainId   = 56
	PolyChainId  = 137
	FTMChainId   = 250
	TPLUSChainId = 1122

	DefaultGasLimit = 3000000
)

type ClientConfig struct {
	ChainId              uint64
	RawChainID           string
	Key                  string
	AccountPrefix        string
	Endpoint             string
	BridgeAddr           string
	VaultBridgeAddr      string
	PegBridgeAddr        string
	MsgPackage           string
	KeyringBackend       string
	TransactorPassphrase string
	GasAdjustment        float64
	GasPrices            string
	Timeout              string
	HomeDir              string
}
