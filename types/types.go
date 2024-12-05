package types

import (
	"time"
)

type Empty struct{}

func (e *Empty) UnmarshalJSON(_ []byte) error {
	return nil
}

func (e *Empty) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

// Result result result
type Result struct {
	Status string `json:"status"`
}

// ID is match engine id type
type ID = uint64

// UserID is match engine userid type
type UserID = uint32

// Timestamp is timestamp
type Timestamp = int64

// Time is match engine return time type
// example: 1678774012.607677
type Time float64

func (t Time) String() string {
	return t.Time().String()
}

func (t Time) Second() int64 {
	return int64(t)
}

func (t Time) Millisecond() int64 {
	return int64(t * 1e3)
}

func (t Time) Microsecond() int64 {
	return int64(t * 1e6)
}

func (t Time) Time() time.Time {
	return time.UnixMicro(t.Microsecond())
}

// TimeMillisecond is match engine return millisecond time type
// example: 1678774829483
type TimeMillisecond = int64

// Side is trade direction, 1: Sell(short) 2: Buy(long)
type Side int

const (
	SideAll Side = iota
	Sell         // short
	Buy          // long
)

func (s Side) String() string {
	switch s {
	case Sell:
		return "sell"
	case Buy:
		return "buy"
	}

	return "undefined"
}

// OrderType
//
// 1: Limit
// 2: Market
type OrderType int

const (
	OrderLimit OrderType = iota + 1
	OrderMarket
)

func (t OrderType) IsLimit() bool {
	return t == OrderLimit
}

func (t OrderType) IsMarket() bool {
	return t == OrderMarket
}

func (t OrderType) String() string {
	switch t {
	case OrderLimit:
		return "limit"
	case OrderMarket:
		return "market"
	}

	return "undefined"
}

// TradeRole
//
// 1: Maker
// 2: Taker
type TradeRole int

const (
	RoleMaker TradeRole = iota + 1
	RoleTaker
)

func (r TradeRole) IsMaker() bool {
	return r == RoleMaker
}

func (r TradeRole) IsTaker() bool {
	return r == RoleTaker
}

func (r TradeRole) String() string {
	switch r {
	case RoleMaker:
		return "maker"
	case RoleTaker:
		return "taker"
	}

	return "undefined"
}
