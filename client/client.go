package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/ashexchange/kit/jsonx"
	"github.com/ashexchange/matchclient/v2/types"
	"github.com/valyala/bytebufferpool"
)

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}

// MaxConnections 最大连接数量，超出该数量的连接会 Block，直到其他请求释放连接。
const MaxConnections = 1000

var defaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: defaultTransportDialContext(&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}),

	ForceAttemptHTTP2:     true,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,

	MaxIdleConns:        MaxConnections,
	MaxIdleConnsPerHost: MaxConnections,
	MaxConnsPerHost:     MaxConnections,
}

type Client struct {
	endpoint string

	cli *http.Client

	reqPool  bytebufferpool.Pool
	respPool bytebufferpool.Pool

	debug bool
	log   func(message string, args ...any)
}

func NewClient(endpoint string, opts ...Option) *Client {
	options := &Options{
		transport: defaultTransport.Clone(),
		log:       debugf,
	}
	for _, o := range opts {
		o(options)
	}

	c := &Client{
		endpoint: endpoint,
		cli: &http.Client{
			Transport: options.transport,
		},
		debug: options.debug,
		log:   options.log,
	}

	return c
}

func (c *Client) Invoke(ctx context.Context, method string, result any, params ...any) error {
	resp, err := c.invoke(ctx, types.NewRequest(FromContext(ctx), method, params...))
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK && resp.ContentLength == 0 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	return c.parseBody(resp.Body, result)
}

func (c *Client) invoke(ctx context.Context, r *types.Request) (resp *http.Response, err error) {
	buf := c.reqPool.Get()
	defer c.reqPool.Put(buf)
	if err = jsonx.NewEncoder(buf).Encode(r); err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewBuffer(buf.B))
	if err != nil {
		return
	}

	if c.debug {
		start := time.Now()

		if resp, err = c.cli.Do(req); err == nil {
			c.log("request: id=%d, method=%s, endpoint=%s, request=%s, elapsed=%s",
				r.Id,
				r.Method,
				c.endpoint,
				buf.B[:buf.Len()-1],
				time.Since(start),
			)

			return
		}

		c.log("request: id=%d, method=%s, endpoint=%s, request=%s, elapsed=%s, error=%s",
			r.Id,
			r.Method,
			c.endpoint,
			buf.B[:buf.Len()-1],
			time.Since(start),
			err,
		)

		return
	}

	return c.cli.Do(req)
}

func (c *Client) parseBody(r io.Reader, result any) error {
	buf := c.respPool.Get()
	defer c.respPool.Put(buf)
	if _, err := io.Copy(buf, r); err != nil {
		return err
	}

	if c.debug && buf.Len() < 1024 {
		c.log("response: %s", buf.B)
	}

	if result == nil {
		result = &types.Empty{}
	}

	reply := types.Response{
		Result: result,
	}
	if err := jsonx.Unmarshal(buf.B, &reply); err != nil {
		return fmt.Errorf("jsonx Unmarshal: %w", err)
	}

	if err := reply.Error; err != nil {
		return err
	}

	return nil
}
