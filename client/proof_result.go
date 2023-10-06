package client

type ProofResultData struct {
	Msg       string `json:"message,omitempty"`
	Sig       []byte `json:"signature,omitempty"`
	RoundId   uint64 `json:"round_id,omitempty"`
	PublicKey []byte `json:"public_key,omitempty"`
	PubKeyStr string `json:"pub_key_str,omitempty"`
}

func (p *ProofResultData) Message() string {
	return p.Msg
}

func (p *ProofResultData) Signature() []byte {
	return p.Sig
}

func (p *ProofResultData) Round() uint64 {
	return p.RoundId
}

func (p *ProofResultData) PubKey() []byte {
	return p.PublicKey
}

func (p *ProofResultData) PubKeyOrg() string {
	return p.PubKeyStr
}
