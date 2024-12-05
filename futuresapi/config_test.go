package futuresapi

import (
	"testing"
)

var cc = NewConfigClient(cli)

func TestConfigClient_UpdateAsset(t *testing.T) {
	if err := cc.UpdateAsset(ctx); err != nil {
		t.Error(err)
	}
}

func TestConfigClient_UpdateIndex(t *testing.T) {
	if err := cc.UpdateIndex(ctx); err != nil {
		t.Error(err)
	}
}

func TestConfigClient_UpdateMarket(t *testing.T) {
	if err := cc.UpdateMarket(ctx); err != nil {
		t.Error(err)
	}
}
