package utils

import "github.com/drand/drand/verify/proto"

func CalculateMintId(mint *proto.Mint) ([32]byte, error) {
	return [32]byte{}, nil
}

func CalculateWithdrawId(withdraw *proto.Withdraw) ([32]byte, error) {
	return [32]byte{}, nil
}

func EncodeMsgSign(domain []byte, msgRaw []byte) ([]byte, error) {
	return []byte{}, nil
}

func GetDomain(addr []byte, action string, chainId uint64) ([]byte, error) {
	return []byte{}, nil
}
