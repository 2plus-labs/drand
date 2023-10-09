package cosmos

import (
	"encoding/binary"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/drand/drand/log"
	"github.com/drand/drand/verify/config"
	"github.com/drand/drand/verify/proto"
	"github.com/drand/drand/verify/utils"
	"github.com/ethereum/go-ethereum/common"
	ec "github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"net/http"
)

var IsTest = false

type CoClient struct {
	client     *http.Client
	endPoint   string
	ChainID    uint64
	RawChainID string // the original cosmos chain id

	BridgeAddr    string // cosmos canonical hex addr
	VaultAddr     string // cosmos canonical hex addr
	PegBridgeAddr string // cosmos canonical hex addr

	logger log.Logger
}

func NewCoClient(cfg config.ChainInfo, logger log.Logger) (*CoClient, error) {
	client := &http.Client{}
	ret := &CoClient{
		client:        client,
		endPoint:      cfg.RPC,
		ChainID:       cfg.ChainId,
		BridgeAddr:    cfg.BridgeAddr,
		VaultAddr:     cfg.VaultBridgeAddr,
		PegBridgeAddr: cfg.PegBridgeAddr,
		logger:        logger,
	}

	return ret, nil
}

func (c *CoClient) GetDomainTokenVault(action string) ([32]byte, error) {
	accAddr, err := c.GetContractAddrBytes(c.VaultAddr)
	if err != nil {
		return [32]byte{}, err
	}

	return utils.GetDomain(accAddr, action, c.ChainID)
}

func (c *CoClient) GetDomainPeggedToken(action string) ([32]byte, error) {
	accAddr, err := c.GetContractAddrBytes(c.PegBridgeAddr)
	if err != nil {
		return [32]byte{}, err
	}

	return utils.GetDomain(accAddr, action, c.ChainID)
}

func (c *CoClient) GetMintDestChainId(mint *proto.Mint) (uint64, error) {
	if IsTest {
		return config.BSCChainId, nil
	}

	// checking have deposit into vault by refId
	vaultRecord, err := c.QueryVaultRecord(common.HexToHash(string(mint.RefId)))
	if err != nil {
		c.logger.Errorw("Failed to query vault record", "error", err)
		return 0, err
	}
	if vaultRecord == 0 {
		c.logger.Errorw("Deposit record not found", "ref_id", mint.RefId)
		return 0, fmt.Errorf("deposit record not found")
	}

	return vaultRecord, nil
}

func (c *CoClient) VerifyMintMsgOnDest(mint *proto.Mint) error {
	if IsTest {
		return nil
	}

	// calculate mint id and check if it existed
	mintId, err := utils.CalculateMintIdV1(mint)
	if err != nil {
		c.logger.Errorw("Failed to calculate mint id", "error", err)
		return err
	}
	mintRecord, err := c.QueryPegBridgeRecord(mintId)
	if err != nil {
		c.logger.Errorw("Failed to query peg bridge record", "error", err)
		return err
	}
	if mintRecord != 0 {
		c.logger.Errorw("Mint record existed", "mint_id", mintId)
		return fmt.Errorf("mint record existed")
	}

	return nil
}

func (c *CoClient) GetWithdrawDestChainId(withdraw *proto.Withdraw) (uint64, error) {
	// checking have burn on pegged smart contract by refId
	bridgeRecord, err := c.QueryPegBridgeRecord(common.HexToHash(string(withdraw.RefId)))
	if err != nil {
		c.logger.Errorw("Failed to query peg bridge record", "error", err)
		return 0, err
	}
	if bridgeRecord == 0 {
		c.logger.Errorw("Withdraw record not found", "ref_id", withdraw.RefId)
		return 0, fmt.Errorf("withdraw record not found")
	}

	return bridgeRecord, nil
}

func (c *CoClient) VerifyWithdrawMsgOnDest(withdraw *proto.Withdraw) error {
	// calculate withdraw id and check if it existed
	withdrawId, err := utils.CalculateWithdrawIdV1(withdraw)
	if err != nil {
		c.logger.Errorw("Failed to calculate withdraw id", "error", err)
		return err
	}
	vaultRecord, err := c.QueryVaultRecord(withdrawId)
	if err != nil {
		c.logger.Errorw("Failed to query vault record", "error", err)
		return err
	}
	if vaultRecord != 0 {
		c.logger.Errorw("Withdraw record existed", "withdraw_id", withdrawId)
		return fmt.Errorf("withdraw record existed")
	}

	return nil
}

func (c *CoClient) QueryVaultRecord(id ec.Hash) (uint64, error) {
	url := fmt.Sprintf("%s/wasm/contract/%s/smart/%s", c.endPoint, c.VaultAddr, fmt.Sprintf("{\"record\":{\"id\":\"%x\"}}", id))
	resp, err := c.Get(url)
	if err != nil {
		c.logger.Errorw("Failed to query vault record", "error", err)
		return 0, err
	}
	c.logger.Info("Query vault record", "resp", resp)

	return binary.BigEndian.Uint64([]byte(resp)), nil
}

func (c *CoClient) QueryPegBridgeRecord(id ec.Hash) (uint64, error) {
	url := fmt.Sprintf("%s/wasm/contract/%s/smart/%s", c.endPoint, c.PegBridgeAddr, fmt.Sprintf("{\"record\":{\"id\":\"%x\"}}", id))
	resp, err := c.Get(url)
	if err != nil {
		c.logger.Errorw("Failed to query peg bridge record", "error", err)
		return 0, err
	}
	c.logger.Info("Query peg bridge record", "resp", resp)

	return binary.BigEndian.Uint64([]byte(resp)), nil
}

func (c *CoClient) GetContractAddrBytes(contractCanonicalAddr string) ([]byte, error) {
	accAddr, err := sdk.AccAddressFromBech32(contractCanonicalAddr)
	if err != nil {
		return nil, err
	}

	return accAddr.Bytes(), nil
}

func (c *CoClient) Get(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	c.logger.Info("Get http response", "url", url, "status", resp.Status, "body", string(body))

	return string(body), nil
}
