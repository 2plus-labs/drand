package evm

import (
	"context"
	"fmt"
	"github.com/drand/drand/log"
	"github.com/drand/drand/verify/abi/peggedtoken"
	"github.com/drand/drand/verify/abi/tokenvault"
	"github.com/drand/drand/verify/config"
	"github.com/drand/drand/verify/proto"
	"github.com/drand/drand/verify/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	IsTest = false
)

type EthClient struct {
	client     *ethclient.Client
	ChainID    uint64
	RawChainID string // the original cosmos chain id

	BridgeAddr    string // eth canonical hex addr
	VaultAddr     string // eth canonical hex addr
	PegBridgeAddr string // eth canonical hex addr

	logger              log.Logger
	tokenVaultInstance  *tokenvault.Contracts
	peggedTokenInstance *peggedtoken.Contracts
}

func NewEthClient(cfg config.ChainInfo, logger log.Logger) (*EthClient, error) {
	client, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, err
	}

	tokenVaultAddr := common.HexToAddress(cfg.VaultBridgeAddr)
	tokenVault, err := tokenvault.NewContracts(tokenVaultAddr, client)
	if err != nil {
		return nil, err
	}

	peggedTokenAddr := common.HexToAddress(cfg.PegBridgeAddr)
	peggedToken, err := peggedtoken.NewContracts(peggedTokenAddr, client)
	if err != nil {
		return nil, err
	}

	return &EthClient{
		client:              client,
		ChainID:             cfg.ChainId,
		BridgeAddr:          cfg.BridgeAddr,
		VaultAddr:           cfg.VaultBridgeAddr,
		PegBridgeAddr:       cfg.PegBridgeAddr,
		logger:              logger,
		tokenVaultInstance:  tokenVault,
		peggedTokenInstance: peggedToken,
	}, nil
}

func (c *EthClient) Close() {
	c.client.Close()
}

func (c *EthClient) GetDomainTokenVault(action string) ([32]byte, error) {
	vAddr := common.HexToAddress(c.VaultAddr)
	return utils.GetDomain(vAddr.Bytes(), action, c.ChainID)
}

func (c *EthClient) GetDomainPeggedToken(action string) ([32]byte, error) {
	pAddr := common.HexToAddress(c.PegBridgeAddr)
	return utils.GetDomain(pAddr.Bytes(), action, c.ChainID)
}

func (c *EthClient) GetMintDestChainId(mint *proto.Mint) (uint64, error) {
	// use for test sign
	if IsTest {
		return config.TPLUSChainId, nil
	}

	// checking have deposit into vault by refId
	depId := [32]byte{}
	copy(depId[:], mint.RefId[:32])
	destChainId, err := c.tokenVaultInstance.Records(&bind.CallOpts{Context: context.Background()}, depId)
	if err != nil {
		c.logger.Errorw("Failed to call contract token vault get record", "error", err)
		return 0, err
	}
	temp := destChainId.String()
	c.logger.Debugw("GetMintDestChainId", "refId", mint.RefId, "destChainId", temp)
	if destChainId.Int64() == 0 {
		c.logger.Errorw("deposit id not existed", "refId", mint.RefId)
		return 0, fmt.Errorf("deposit id not existed")
	}

	return destChainId.Uint64(), nil
}

// VerifyMintMsgOnDest verify tx from src chain exited (deposit tx into vault)
func (c *EthClient) VerifyMintMsgOnDest(mint *proto.Mint) error {
	if IsTest {
		return nil
	}

	// calculate mintId and check mintId is existed or not on the destination chain
	mintId, err := utils.CalculateMintIdV2(mint, common.HexToAddress(c.PegBridgeAddr).Bytes())
	if err != nil {
		c.logger.Errorw("Failed to calculate mint id", "error", err)
		return err
	}
	mintIdExist, err := c.peggedTokenInstance.Records(&bind.CallOpts{Context: context.Background()}, mintId)
	if err != nil {
		c.logger.Errorw("Failed to call contract pegged token get record", "error", err)
		return err
	}
	if mintIdExist.Int64() != 0 {
		c.logger.Errorw("mint id is existed", "mintId", mintId)
		return fmt.Errorf("mint id is existed")
	}

	return nil
}

func (c *EthClient) GetWithdrawDestChainId(withdraw *proto.Withdraw) (uint64, error) {
	// checking have burn on pegged smart contract by refId
	burnId := [32]byte{}
	copy(burnId[:], withdraw.RefId[:32])
	destChainId, err := c.peggedTokenInstance.Records(&bind.CallOpts{Context: context.Background()}, burnId)
	if err != nil {
		c.logger.Errorw("Failed to call contract pegged token get record", "error", err)
		return 0, err
	}
	if destChainId.Int64() == 0 {
		c.logger.Errorw("burn id not existed", "refId", withdraw.RefId)
		return 0, fmt.Errorf("burn id not existed")
	}

	return destChainId.Uint64(), nil
}

// VerifyWithdrawMsgOnDest verify msg withdraw do not use before
func (c *EthClient) VerifyWithdrawMsgOnDest(withdraw *proto.Withdraw) error {
	// calculate withdrawId and check withdrawId is existed or not on the destination chain
	withdrawId, err := utils.CalculateWithdrawIdV2(withdraw, common.HexToAddress(c.VaultAddr).Bytes())
	if err != nil {
		c.logger.Errorw("Failed to calculate withdraw id", "error", err)
		return err
	}
	withdrawIdExist, err := c.tokenVaultInstance.Records(&bind.CallOpts{Context: context.Background()}, withdrawId)
	if err != nil {
		c.logger.Errorw("Failed to call contract token vault get record", "error", err)
		return err
	}
	if withdrawIdExist.Int64() != 0 {
		c.logger.Errorw("withdraw id is existed", "withdrawId", withdrawId)
		return fmt.Errorf("withdraw id is existed")
	}

	return nil
}
