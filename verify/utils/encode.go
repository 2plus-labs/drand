package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/drand/drand/verify/proto"
	"github.com/ethereum/go-ethereum/crypto"
)

func CalculateMintIdV1(mint *proto.Mint) ([32]byte, error) {
	refChainId := make([]byte, 8)
	binary.BigEndian.PutUint64(refChainId, mint.RefChainId)
	var amount [32]byte
	from := 32 - len(mint.Amount)
	copy(amount[from:], mint.Amount)
	dataHash := encodePacked(
		mint.Account,
		mint.Token,
		amount[:],
		mint.Depositor,
		refChainId,
		mint.RefId,
	)

	return crypto.Keccak256Hash(dataHash), nil
}

func CalculateMintIdV2(mint *proto.Mint, peggedAddr []byte) ([32]byte, error) {
	refChainId := make([]byte, 8)
	binary.BigEndian.PutUint64(refChainId, mint.RefChainId)
	var amount [32]byte
	from := 32 - len(mint.Amount)
	copy(amount[from:], mint.Amount)
	dataHash := encodePacked(
		mint.Account,
		mint.Token,
		amount[:],
		mint.Depositor,
		refChainId,
		mint.RefId,
		peggedAddr,
	)

	return crypto.Keccak256Hash(dataHash), nil
}

func CalculateWithdrawIdV1(withdraw *proto.Withdraw) ([32]byte, error) {
	refChainId := make([]byte, 8)
	binary.BigEndian.PutUint64(refChainId, withdraw.RefChainId)
	var amount [32]byte
	from := 32 - len(withdraw.Amount)
	copy(amount[from:], withdraw.Amount)
	dataHash := encodePacked(
		withdraw.Receiver,
		withdraw.Token,
		amount[:],
		withdraw.BurnAccount,
		refChainId,
		withdraw.RefId,
	)
	return crypto.Keccak256Hash(dataHash), nil
}

func CalculateWithdrawIdV2(withdraw *proto.Withdraw, tokenVaultAddr []byte) ([32]byte, error) {
	refChainId := make([]byte, 8)
	binary.BigEndian.PutUint64(refChainId, withdraw.RefChainId)
	var amount [32]byte
	from := 32 - len(withdraw.Amount)
	copy(amount[from:], withdraw.Amount)
	dataHash := encodePacked(
		withdraw.Receiver,
		withdraw.Token,
		amount[:],
		withdraw.BurnAccount,
		refChainId,
		withdraw.RefId,
		tokenVaultAddr,
	)
	return crypto.Keccak256Hash(dataHash), nil
}

func EncodeMsgSign(domain [32]byte, msgRaw []byte) ([]byte, error) {
	return encodePacked(domain[:], msgRaw), nil
}

func GetDomain(addr []byte, action string, chainId uint64) ([32]byte, error) {
	chainIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(chainIdBytes, chainId)
	return crypto.Keccak256Hash(encodePacked(chainIdBytes, addr, []byte(action))), nil
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
