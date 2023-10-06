package evm

import (
	"encoding/hex"
	"github.com/drand/drand/log"
	"github.com/drand/drand/verify/config"
	"github.com/drand/drand/verify/proto"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func newClient() *EthClient {

	//Bridge=0xA73F5eb34679D7b8fA3a9084bE23F5F44aeb07dB
	//
	//OriginalTokenVaultV2=0xb3406163887752a55F8203229720272A50A10a61
	//
	//PeggedTokenBridgeV2=0x9216fC92F4Cf71d4c3Cc5f7A4bA5e742d2890a1c
	ethCfg := config.ChainInfo{
		ChainId: config.BSCChainId,
		//RPC:             "https://data-seed-prebsc-1-s1.binance.org:8545/", // testnet
		RPC:             "https://bsc-dataseed1.binance.org/", // mainnet
		BridgeAddr:      "0xA73F5eb34679D7b8fA3a9084bE23F5F44aeb07dB",
		VaultBridgeAddr: "0x78bc5Ee9F11d133A08b331C2e18fE81BE0Ed02DC",
		PegBridgeAddr:   "0x9216fC92F4Cf71d4c3Cc5f7A4bA5e742d2890a1c",
		Type:            config.EVM,
	}
	logger := log.NewLogger(nil, log.LogDebug).Named("EVMTest")
	evmClient, err := NewEthClient(ethCfg, logger)
	if err != nil {
		panic(err)
	}
	return evmClient
}

func TestEthClient_GetDomainTokenVault(t *testing.T) {
	evmClient := newClient()
	domain, err := evmClient.GetDomainTokenVault("mint")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(domain[:]))
}

func TestEthClient_GetDomainPeggedToken(t *testing.T) {
	evmClient := newClient()
	domain, err := evmClient.GetDomainPeggedToken("withdraw")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(domain[:]))
}

func TestEthClient_GetMintDestChainId(t *testing.T) {
	evmClient := newClient()
	depId, err := hex.DecodeString("804012600E433229699E806F5809221200FD2DA84CB8314D5C303470D9C12983")
	if err != nil {
		t.Fatal(err)
	}
	amount := big.NewInt(10576801887)
	mint := &proto.Mint{
		Token:      common.HexToAddress("0x7ad7242A99F21aa543F9650A56D141C57e4F6081").Bytes(),
		Account:    common.HexToAddress("0x88840c37e8fa21afe5c28F0c704F9dcbd1601a45").Bytes(),
		Amount:     amount.Bytes(),
		Depositor:  common.HexToAddress("0x88840c37e8fa21afe5c28F0c704F9dcbd1601a45").Bytes(),
		RefChainId: config.BSCChainId,
		RefId:      depId,
	}

	chainId, err := evmClient.GetMintDestChainId(mint)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(chainId)
}
