package futuresapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2/types"
)

type MonitorClient struct {
	invoker Invoker
}

func NewMonitorClient(invoker Invoker) MonitorClient {
	return MonitorClient{invoker}
}

type MonitorListScopeResult []string

func (c MonitorClient) ListScope(ctx context.Context) (result MonitorListScopeResult, err error) {
	err = c.invoker.Invoke(ctx, "monitor.list_scope", &result)

	return
}

type MonitorListKeyRequest struct {
	Scope string
}

type MonitorListKeyResult []string

func (c MonitorClient) ListKey(ctx context.Context, req MonitorListKeyRequest) (result MonitorListKeyResult, err error) {
	err = c.invoker.Invoke(ctx, "monitor.list_key", &result, req.Scope)

	return
}

type MonitorListHostRequest struct {
	Scope string
	Key   string
}

type MonitorListHostResult []string

func (c MonitorClient) ListHost(ctx context.Context, req MonitorListHostRequest) (result MonitorListHostResult, err error) {
	err = c.invoker.Invoke(ctx, "monitor.list_host", &result,
		req.Scope,
		req.Key,
	)

	return
}

type MonitorQueryMinuteRequest struct {
	Scope  string
	Key    string
	Host   string
	Points uint32
}

// MonitorCounter
//
// [<timestamp>, <count>]
//
// example: [1681430400, 0]
type MonitorCounter [2]int64

func (item MonitorCounter) Timestamp() int64 {
	return item[0]
}

func (item MonitorCounter) Time() types.Time {
	return types.Time(item.Timestamp())
}

func (item MonitorCounter) Count() int64 {
	return item[1]
}

type MonitorQueryMinuteResult []MonitorCounter

func (c MonitorClient) QueryMinute(ctx context.Context, req MonitorQueryMinuteRequest) (result MonitorQueryMinuteResult, err error) {
	err = c.invoker.Invoke(ctx, "monitor.query_minute", &result,
		req.Scope,
		req.Key,
		req.Host,
		req.Points,
	)

	return
}

type MonitorQueryDailyRequest struct {
	Scope  string
	Key    string
	Host   string
	Points uint32
}

type MonitorQueryDailyResult []MonitorCounter

func (c MonitorClient) QueryDaily(ctx context.Context, req MonitorQueryDailyRequest) (result MonitorQueryDailyResult, err error) {
	err = c.invoker.Invoke(ctx, "monitor.query_daily", &result,
		req.Scope,
		req.Key,
		req.Host,
		req.Points,
	)

	return
}
