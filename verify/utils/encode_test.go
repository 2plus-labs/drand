package utils

import (
	"encoding/hex"
	"github.com/drand/drand/verify/config"
	"github.com/drand/drand/verify/proto"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestCalculateMintId(t *testing.T) {
	depId, err := hex.DecodeString("804012600E433229699E806F5809221200FD2DA84CB8314D5C303470D9C12983")
	if err != nil {
		t.Fatal(err)
	}
	//amount := big.NewInt(10468717106339000000)
	amount := new(big.Int)
	amount.SetString("10468717106339000000", 10)
	mint := &proto.Mint{
		//Token:      common.HexToAddress("0x7ad7242A99F21aa543F9650A56D141C57e4F6081").Bytes(),
		Token:      common.HexToAddress("0x80b010450fdaf6a3f8df033ee296e92751d603b3").Bytes(),
		Account:    common.HexToAddress("0x88840c37e8fa21afe5c28F0c704F9dcbd1601a45").Bytes(),
		Amount:     amount.Bytes(),
		Depositor:  common.HexToAddress("0x88840c37e8fa21afe5c28F0c704F9dcbd1601a45").Bytes(),
		RefChainId: config.BSCChainId,
		RefId:      depId,
	}

	mintId, err := CalculateMintIdV1(mint)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hex.EncodeToString(mintId[:]))
}
