package proto

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/drand/drand/verify/config"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestMintMsg(t *testing.T) {
	mintMsg := &Mint{
		Token:      common.HexToAddress("").Bytes(),
		Account:    common.HexToAddress("").Bytes(),
		Amount:     []byte("100"),
		Depositor:  common.HexToAddress("").Bytes(),
		RefChainId: config.TPLUSChainId,
		RefId:      []byte("1234567890"),
	}

	data, err := json.Marshal(mintMsg)
	if err != nil {
		t.Fatal(err)
	}
	encodeData := base64.StdEncoding.EncodeToString(data)
	t.Log(encodeData)

	decodeData, err := base64.StdEncoding.DecodeString(encodeData)
	if err != nil {
		t.Fatal(err)
	}

	var mint Mint
	err = json.Unmarshal(decodeData, &mint)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mint.String())

	decode2, err := hex.DecodeString("7b22746f6b656e223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030222c226163636f756e74223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030222c22616d6f756e74223a22333133303330222c226465706f7369746f72223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030222c227265665f636861696e5f6964223a313132322c227265665f6964223a223331333233333334333533363337333833393330227d")
	if err != nil {
		t.Fatal(err)
	}

	var mint2 Mint
	err = json.Unmarshal(decode2, &mint2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(mint2.String())
}
