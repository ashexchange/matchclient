package spotapi

import (
	"context"

	"github.com/ashexchange/matchclient/v2/client"
)

const (
	endpoint = "http://landry-ces-accesshttp.aie.test"
	userId   = 1001
)

var (
	ctx = context.Background()
	cli = client.NewClient(endpoint, client.WithDebug(true))
)
