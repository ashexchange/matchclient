package futuresapi

import "github.com/ashexchange/matchclient/v2/types"

// DealType
//
//	1 open position
//	2 add position
//	3 sub position
//	4 close position
//	5 sys close
//	6 position liq
//	7 position adl
type DealType int

const (
	DealOpenPosition DealType = iota + 1
	DealAddPosition
	DealSubPosition
	DealClosePosition
	DealSysClose
	DealPositionLiq
	DealPositionAdl
)

func (t DealType) String() string {
	switch t {
	case DealOpenPosition:
		return "open"
	case DealAddPosition:
		return "add"
	case DealSubPosition:
		return "sub"
	case DealClosePosition:
		return "close"
	case DealSysClose:
		return "sys-close"
	case DealPositionLiq:
		return "liq"
	case DealPositionAdl:
		return "adl"
	}

	return "undefined"
}

// FinishType
// 1: normal
// 2: liq
// 3: adl
type FinishType int

const (
	FinishNormal FinishType = iota + 1
	FinishLiq
	FinishAdl
)

func (ft FinishType) String() string {
	switch ft {
	case FinishNormal:
		return "normal"
	case FinishLiq:
		return "liq"
	case FinishAdl:
		return "adl"
	}

	return "unknown"
}

func (ft FinishType) IsNormal() bool { return ft == FinishNormal }
func (ft FinishType) IsLiq() bool    { return ft == FinishLiq }
func (ft FinishType) IsAdl() bool    { return ft == FinishAdl }

type MarketType int

const (
	Forward MarketType = iota + 1
	Reverse
)

func (m MarketType) String() string {
	switch m {
	case Forward:
		return "forward"
	case Reverse:
		return "reverse"
	}

	return "undefined"
}

func (m MarketType) IsForward() bool { return m == Forward }
func (m MarketType) IsReverse() bool { return m == Reverse }

// OrderEffectType
//
// 1: default, good util cancel
// 2: immediate or cancel
// 3: fill or kill
type OrderEffectType uint32

const (
	EffectDefault OrderEffectType = iota + 1
	EffectImmediateOrCancel
	EffectFillOrKill
)

func (et OrderEffectType) String() string {
	switch et {
	case EffectDefault:
		return "default"
	case EffectImmediateOrCancel:
		return "immediate-or-cancel"
	case EffectFillOrKill:
		return "fill-or-kill"
	}

	return "undefined"
}

func (et OrderEffectType) IsValid() bool {
	switch et {
	case EffectDefault, EffectImmediateOrCancel, EffectFillOrKill:
		return true
	default:
		return false
	}
}

// OrderOption
//
// 1: only maker
// 2: hidden order
type OrderOption uint32

const (
	OptionOnlyMaker OrderOption = iota + 1
	OptionHiddenOrder
)

func (o OrderOption) String() string {
	switch o {
	case OptionOnlyMaker:
		return "only-maker"
	case OptionHiddenOrder:
		return "hidden-order"
	}

	return "undefined"
}

func (o OrderOption) IsValid() bool {
	switch o {
	case OptionOnlyMaker, OptionHiddenOrder:
		return true
	default:
		return false
	}
}

// PositionAdjustMarginType
// 1: add margin
// 2: sub margin
type PositionAdjustMarginType = int

const (
	MarginAdd PositionAdjustMarginType = iota + 1
	MarginSub
)

// PositionType
//
// 1: isolate
// 2: cross
type PositionType int

const (
	Isolated PositionType = iota + 1
	Cross
)

func (p PositionType) IsIsolated() bool { return p == Isolated }
func (p PositionType) IsCross() bool    { return p == Cross }
func (p PositionType) IsValid() bool    { return p.IsIsolated() || p.IsCross() }

func (p PositionType) String() string {
	switch p {
	case Isolated:
		return "isolated"
	case Cross:
		return "cross"
	}

	return "undefined"
}

// Target
//
// 0: undecided
// 1: open position
// 2: add position
// 3: sub position
// 4: close position
// 5: sys close
// 6: position liq
// 7: position adl
type Target int

const (
	TargetUndecided Target = iota
	TargetOpenPosition
	TargetAddPosition
	TargetSubPosition
	TargetClosePosition
	TargetSysClose
	TargetPositionLiq
	TargetPositionAdl
)

func (t Target) String() string {
	switch t {
	case TargetUndecided:
		return "undecided"
	case TargetOpenPosition:
		return "open"
	case TargetAddPosition:
		return "add"
	case TargetSubPosition:
		return "sub"
	case TargetClosePosition:
		return "close"
	case TargetSysClose:
		return "sys-close"
	case TargetPositionLiq:
		return "liq"
	case TargetPositionAdl:
		return "adl"
	}

	return "undefined"
}

// BrkDirection is liquidation price break direction
type BrkDirection int

const (
	BrkDirectionUp   BrkDirection = 1
	BrkDirectionDown BrkDirection = -1
)

func (d BrkDirection) String() string {
	switch d {
	case BrkDirectionUp:
		return "up"
	case BrkDirectionDown:
		return "down"
	}

	return "undefined"
}

func (d BrkDirection) IsUp() bool   { return d == BrkDirectionUp }
func (d BrkDirection) IsDown() bool { return d == BrkDirectionDown }

type PositionMode int

const (
	// OneWay 单仓
	OneWay PositionMode = iota + 1
	// Hedge 多仓
	Hedge
)

func (m PositionMode) String() string {
	switch m {
	case OneWay:
		return "one-way"
	case Hedge:
		return "hedge"
	}

	return "undefined"
}

func (m PositionMode) IsOneWay() bool {
	return m == OneWay
}

func (m PositionMode) IsHedge() bool {
	return m == Hedge
}

func (m PositionMode) IsValid() bool {
	return m.IsOneWay() || m.IsHedge()
}

// Direction
// 1: open
// 2: close
type Direction int

const (
	DirectionOpen Direction = iota + 1
	DirectionClose
)

func (d Direction) String() string {
	switch d {
	case DirectionOpen:
		return "open"
	case DirectionClose:
		return "close"
	}

	return "undefined"
}

func (d Direction) IsOpen() bool  { return d == DirectionOpen }
func (d Direction) IsClose() bool { return d == DirectionClose }

var (
	Short = types.Sell
	Long  = types.Buy
)
