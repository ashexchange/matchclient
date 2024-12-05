package client

import (
	"net/http"
	"time"
)

type Options struct {
	transport *http.Transport
	debug     bool
	log       func(message string, args ...any)
}

type Option func(c *Options)

func WithTransport(t *http.Transport) Option {
	return func(c *Options) {
		c.transport = t
	}
}

func WithMaxConnections(n int) Option {
	return func(c *Options) {
		c.transport.MaxIdleConns = n
		c.transport.MaxIdleConnsPerHost = n
		c.transport.MaxConnsPerHost = n
	}
}

func WithIdleConnTimeout(timeout time.Duration) Option {
	return func(c *Options) {
		c.transport.IdleConnTimeout = timeout
	}
}

func WithDebug(debug bool) Option {
	return func(c *Options) {
		c.debug = debug
	}
}

func WithLog(log func(message string, args ...any)) Option {
	return func(c *Options) {
		c.log = log
	}
}
