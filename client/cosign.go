package client

// CoSignData hold co-sign data of multi-party
type CoSignData struct {
	Msg       string `json:"message,omitempty"`
	Sig       []byte `json:"signature,omitempty"`
	Random    []byte `json:"randomness,omitempty"`
	PublicKey []byte `json:"public_key"`
}

func (c *CoSignData) PubKey() []byte {
	return c.PublicKey
}

func (c *CoSignData) Randomness() []byte {
	return c.Random
}

func (c *CoSignData) Signature() []byte {
	return c.Sig
}

func (c *CoSignData) Message() string {
	return c.Msg
}
