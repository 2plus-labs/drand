package config

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewCfg(t *testing.T) {
	cfgClients, err := NewCfgClient("/Users/yeshidolma/working_space/2plus/drand/verify/config")
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(cfgClients)
}
