package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

const (
	EthChainId   = 1
	BSCChainId   = 56
	PolyChainId  = 137
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
		RPC:             "https://bsc-dataseed.binance.org",
		BridgeAddr:      "0x4C6Aa7E4788aF957E3FEa0B6eDc8E28b1f8a7d8f",
		VaultBridgeAddr: "0x4C6Aa7E4788aF957E3FEa0B6eDc8E28b1f8a7d8f",
		PegBridgeAddr:   "0x4C6Aa7E4788aF957E3FEa0B6eDc8E28b1f8a7d8f",
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
		RPC:             "https://rpc-mainnet.maticvigil.com",
		BridgeAddr:      "",
		VaultBridgeAddr: "",
		PegBridgeAddr:   "",
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
