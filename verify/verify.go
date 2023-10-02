package verify

import (
	"github.com/drand/drand/log"
	"github.com/drand/drand/verify/clients/cosmos"
	"github.com/drand/drand/verify/clients/evm"
	"github.com/drand/drand/verify/config"
	"github.com/drand/drand/verify/proto"
)

type BridgeMsg interface {
	// GetMintDestChainId returns the destination chain id of the mint message
	// it will be checking have deposit into vault by refId
	// if not exist, return error
	GetMintDestChainId(mint *proto.Mint) (uint64, error)

	// VerifyMintMsgOnDest returns nil if the mint message is not exist on the destination chain
	// It means the mint message able to use to mint token on the destination chain
	// It will be return error if the mint message is existed on the destination chain
	VerifyMintMsgOnDest(mint *proto.Mint) error

	// GetDomainTokenVault returns the domain of the message
	GetDomainTokenVault(action string) ([]byte, error)

	// GetDomainPeggedToken returns the domain of the message
	GetDomainPeggedToken(action string) ([]byte, error)

	// GetWithdrawDestChainId returns the destination chain id of the withdraw message
	// it will be checking have burn on pegged smart contract by refId
	// if not exist, return error
	GetWithdrawDestChainId(withdraw *proto.Withdraw) (uint64, error)

	// VerifyWithdrawMsgOnDest returns nil if the withdraw message is not exist on the destination chain
	// It means the withdraw message able to use to burn token on the destination chain
	// It will be return error if the withdraw message is existed on the destination chain
	VerifyWithdrawMsgOnDest(withdraw *proto.Withdraw) error
}

type Proxy struct {
	chainSupported map[uint64]BridgeMsg
}

func NewVerifyProxy(cfgClient *config.CfgClient, logger log.Logger) (*Proxy, error) {
	chains := make(map[uint64]BridgeMsg)
	for _, chainInfo := range cfgClient.Chains {
		if chainInfo.Type == config.EVM {
			evmClient, err := evm.NewEthClient(chainInfo, logger)
			if err != nil {
				return nil, err
			}
			chains[chainInfo.ChainId] = evmClient
		} else {
			cosClient, err := cosmos.NewCoClient(chainInfo, logger)
			if err != nil {
				return nil, err
			}
			chains[chainInfo.ChainId] = cosClient
		}
	}

	return &Proxy{
		chainSupported: chains,
	}, nil
}

func (p *Proxy) GetInstance(chainId uint64) BridgeMsg {
	chainClient, ok := p.chainSupported[chainId]
	if !ok {
		return nil
	}
	return chainClient
}
