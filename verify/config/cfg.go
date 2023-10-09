package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

const (
	EthChainId = 1
	//BSCChainId   = 56
	BSCChainId = 97
	//PolyChainId  = 137
	PolyChainId  = 80001
	TPLUSChainId = 1122

	DefaultGasLimit = 3000000
)

type ChainType string

const (
	EVM    ChainType = "EVM"
	COSMOS ChainType = "COSMOS"
)

type ChainInfo struct {
	ChainId         uint64
	RPC             string
	BridgeAddr      string
	VaultBridgeAddr string
	PegBridgeAddr   string
	Type            ChainType
}

type CfgClient struct {
	Chains []ChainInfo
}

func NewCfgClient(cfgPath ...string) (*CfgClient, error) {
	if len(cfgPath) > 0 {
		viper.AddConfigPath(".")
		viper.AddConfigPath(cfgPath[0])
		viper.SetConfigName("networks")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading config file, %v", err)
		}

		allSettings := viper.AllSettings()
		chains := make([]ChainInfo, 0, len(allSettings))

		for key, value := range allSettings {
			val := value.(map[string]interface{})

			if val["chain_id"] == nil {
				return nil, fmt.Errorf("chain_id not defined for chain %s", key)
			}

			if val["rpc_url"] == nil {
				return nil, fmt.Errorf("end_point not defined for chain %s", key)
			}

			if val["bridge_addr"] == nil {
				return nil, fmt.Errorf("bridge_addr not defined for chain %s", key)
			}

			if val["vault_bridge_addr"] == nil {
				return nil, fmt.Errorf("vault_bridge_addr not defined for chain %s", key)
			}

			if val["peg_bridge_addr"] == nil {
				return nil, fmt.Errorf("peg_bridge_addr not defined for chain %s", key)
			}

			if val["type"] == nil {
				return nil, fmt.Errorf("type not defined for chain %s", key)
			}

			chainId, err := strconv.ParseUint(val["chain_id"].(string), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid chain_id for chain %s", key)
			}

			chains = append(chains, ChainInfo{
				ChainId:         chainId,
				RPC:             val["rpc_url"].(string),
				BridgeAddr:      val["bridge_addr"].(string),
				VaultBridgeAddr: val["vault_bridge_addr"].(string),
				PegBridgeAddr:   val["peg_bridge_addr"].(string),
				Type:            ChainType(val["type"].(string)),
			})
		}

		return &CfgClient{Chains: chains}, nil
	}

	return &CfgClient{
		Chains: []ChainInfo{
			defaultBSCChainInfo(),
			defaultEthChainInfo(),
			defaultPolyChainInfo(),
			defaultTPLUSChainInfo(),
		},
	}, nil
}

func defaultBSCChainInfo() ChainInfo {
	return ChainInfo{
		ChainId:         BSCChainId,
		RPC:             "https://bsc-testnet.nodereal.io/v1/73d5dedd6b0f40ecb167818a2586682b",
		BridgeAddr:      "0xe81eA3FC8B06154D89abD28f4811225a953b2FCD",
		VaultBridgeAddr: "0x6c48912Fead677275CB08101e604Af7789868dfD",
		PegBridgeAddr:   "0x1F83f4E2ef9185EA4c5d1bF516202640f83EC150",
		Type:            EVM,
	}
}

func defaultEthChainInfo() ChainInfo {
	return ChainInfo{
		ChainId:         EthChainId,
		RPC:             "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
		BridgeAddr:      "",
		VaultBridgeAddr: "",
		PegBridgeAddr:   "",
		Type:            EVM,
	}
}

func defaultPolyChainInfo() ChainInfo {
	return ChainInfo{
		ChainId:         PolyChainId,
		RPC:             "https://rpc-mumbai.maticvigil.com",
		BridgeAddr:      "0xf04ed13b38C98dA9369B75C883A2EB37983BE4dD",
		VaultBridgeAddr: "0xaF7cea571544109481B7B82d2ed8ee2Cd3339a1F",
		PegBridgeAddr:   "0x5225cbb2440dbA3Fb84191Ba5dC84e17133488Da",
		Type:            EVM,
	}
}

func defaultTPLUSChainInfo() ChainInfo {
	return ChainInfo{
		ChainId:         TPLUSChainId,
		RPC:             "http://localhost:26657",
		BridgeAddr:      "",
		VaultBridgeAddr: "",
		PegBridgeAddr:   "",
		Type:            COSMOS,
	}
}
