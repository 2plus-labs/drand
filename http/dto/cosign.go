package dto

import json "github.com/nikkolasg/hexjson"

type CoSign struct {
	Msg       string `json:"message"`
	Signature string `json:"signature"`
	Round     uint64 `json:"round"`
}

type AddPubKey struct {
	PublicKey string `json:"public_key"`
}

type ErrorResp struct {
	Err string `json:"err"`
}

func (e ErrorResp) ToString() string {
	errBytes, _ := json.Marshal(e)
	return string(errBytes)
}

type WrapChainInfo struct {
	PublicKey   []byte `json:"public_key"`
	Period      uint32 `json:"period"`
	GenesisTime int64  `json:"genesis_time"`
	Hash        []byte `json:"hash"`
	SchemeID    string `json:"scheme_id"`
}
