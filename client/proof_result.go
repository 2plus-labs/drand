package client

type ProofResultData struct {
	Msg       string `json:"message,omitempty"`
	Sig       []byte `json:"signature,omitempty"`
	RoundId   uint64 `json:"round_id,omitempty"`
	PublicKey []byte `json:"public_key,omitempty"`
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
