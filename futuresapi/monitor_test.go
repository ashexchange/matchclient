package futuresapi

import (
	"testing"
)

var monc = NewMonitorClient(cli)

func TestMonitor_ListScope(t *testing.T) {
	call0(t, monc.ListScope)
}

func TestMonitor_ListKey(t *testing.T) {
	call2(t, monc.ListKey, MonitorListKeyRequest{
		Scope: "perpetual_matchengine",
	})
}

func TestMonitor_ListHost(t *testing.T) {
	call2(t, monc.ListHost, MonitorListHostRequest{
		Scope: "perpetual_matchengine",
		Key:   "cmd_order_depth",
	})
}

func TestMonitor_QueryMinute(t *testing.T) {
	call2(t, monc.QueryMinute, MonitorQueryMinuteRequest{
		Scope:  "perpetual_matchengine",
		Key:    "cmd_order_depth",
		Host:   "matchengine",
		Points: 10,
	})
}

func TestMonitor_QueryDaily(t *testing.T) {
	call2(t, monc.QueryDaily, MonitorQueryDailyRequest{
		Scope:  "perpetual_matchengine",
		Key:    "cmd_order_depth",
		Host:   "matchengine",
		Points: 10,
	})
}
