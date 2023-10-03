package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/drand/drand/verify/proto"
	"github.com/ethereum/go-ethereum/crypto"
)

func CalculateMintId(mint *proto.Mint, peggedAddr []byte) ([32]byte, error) {
	//uint64Ty, err := abi.NewType("uint", "uint64", nil)
	//if err != nil {
	//	return [32]byte{}, err
	//}
	//uint256Ty, err := abi.NewType("uint256", "", nil)
	//if err != nil {
	//	return [32]byte{}, err
	//}
	//bytes32Ty, err := abi.NewType("bytes32", "", nil)
	//if err != nil {
	//	return [32]byte{}, err
	//}
	//addressTy, err := abi.NewType("address", "", nil)
	//if err != nil {
	//	return [32]byte{}, err
	//}
	//arguments := abi.Arguments{
	//	{
	//		Name: "token",
	//		Type: addressTy,
	//	},
	//	{
	//		Name: "account",
	//		Type: addressTy,
	//	},
	//	{
	//		Name: "amount",
	//		Type: uint256Ty,
	//	},
	//	{
	//		Name: "depositor",
	//		Type: addressTy,
	//	},
	//	{
	//		Name: "refChainId",
	//		Type: uint64Ty,
	//	},
	//	{
	//		Name: "refId",
	//		Type: bytes32Ty,
	//	},
	//}
	//
	//bytes, err := arguments.Pack(
	//	common.HexToAddress(mint.Token),
	//	[32]byte{'I', 'D', '1'},
	//	big.NewInt(42),
	//)
	refChainId := make([]byte, 8)
	binary.LittleEndian.PutUint64(refChainId, mint.RefChainId)
	dataHash := encodePacked(
		mint.Account,
		mint.Token,
		mint.Amount,
		mint.Depositor,
		refChainId,
		mint.RefId,
		peggedAddr,
	)

	return crypto.Keccak256Hash(dataHash), nil
}

func CalculateWithdrawId(withdraw *proto.Withdraw, tokenVaultAddr []byte) ([32]byte, error) {
	refChainId := make([]byte, 8)
	binary.LittleEndian.PutUint64(refChainId, withdraw.RefChainId)
	dataHash := encodePacked(
		withdraw.Receiver,
		withdraw.Token,
		withdraw.Amount,
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
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return crypto.Keccak256Hash(encodePacked(chainIdBytes, addr, []byte(action))), nil
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
