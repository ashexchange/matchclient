package futuresapi

import (
	"testing"
)

var ic = NewIndexClient(cli)

func TestIndexClient_Debug(t *testing.T) {
	call1(t, ic.Debug, IndexDebugRequest{
		Market:     "BTCUSDT",
		Debug:      true,
		IndexPrice: "21764.52",
	})
}

func TestIndexClient_List(t *testing.T) {
	call0(t, ic.List)
}

func TestIndexClient_Query(t *testing.T) {
	call2(t, ic.Query, IndexQueryRequest{
		Market: "BTCUSDT",
	})
}
