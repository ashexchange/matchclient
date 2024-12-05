package futuresapi

import (
	"errors"

	"github.com/ashexchange/matchclient/v2/types"
)

var (
	ErrHttpInvalidArgument    = types.NewError(2001, "invalid argument")
	ErrHttpServiceUnavailable = types.NewError(2002, "service unavailable")
	ErrHttpServiceTimeout     = types.NewError(2003, "service timeout")
	ErrHttpUnknownMethod      = types.NewError(2004, "unknown method")
	ErrHttpInternalError      = types.NewError(2007, "internal error")
	ErrHttpResultNull         = types.NewError(2008, "result null")
)

var (
	ErrInvalidArgument    = types.NewError(3001, "invalid argument")
	ErrServiceUnavailable = types.NewError(3002, "service unavailable")
	ErrServiceTimeout     = types.NewError(3003, "service timeout")
	ErrUnknownMethod      = types.NewError(3004, "unknown method")
	ErrInternalError      = types.NewError(3006, "internal error")
	ErrRateLimit          = types.NewError(3007, "rate limit")
)

var (
	ErrMarketNotExists         = types.NewError(3101, "market not exists")
	ErrUserIdNotExists         = types.NewError(3102, "user id not exists")
	ErrOrderNotExists          = types.NewError(3103, "order not exists")
	ErrPositionNotExists       = types.NewError(3105, "position not exists")
	ErrAssetNotExists          = types.NewError(3106, "asset not exists")
	ErrBalanceUpdateRepeated   = types.NewError(3107, "balance update repeated")
	ErrAmountExceed            = types.NewError(3108, "amount exceed limit")
	ErrBalanceNotEnough        = types.NewError(3109, "balance not enough")
	ErrTraderNotEnough         = types.NewError(3110, "trader not enough")
	ErrExceedMaxLimit          = types.NewError(3111, "exceed max limit")
	ErrUserNotMatch            = types.NewError(3112, "user not match")
	ErrInvalidLeverage         = types.NewError(3113, "invalid leverage value")
	ErrOrderPriceInvalid       = types.NewError(3114, "order price invalid")
	ErrPositionIsLiquidation   = types.NewError(3115, "position liquidating")
	ErrCanNotCompleteDeal      = types.NewError(3116, "can not complete deal, kill order")
	ErrPositionWillLiquidation = types.NewError(3117, "position will liquidation")
	ErrLimitPriceLower         = types.NewError(3118, "limit price lower than liquidation price")
	ErrLimitPriceHigher        = types.NewError(3119, "limit price higher than liquidation price")
	ErrNoPosition              = types.NewError(3120, "no position")
	ErrCrossPosition           = types.NewError(3121, "cross position")
	ErrOrderExists             = types.NewError(3122, "order exist")
	ErrCanNotSubMargin         = types.NewError(3123, "margin less init margin")
	ErrSubTooMuchMargin        = types.NewError(3124, "sub too much margin")
	ErrAmountTooSmall          = types.NewError(3127, "amount too small")
	ErrInvalidPriceSize        = types.NewError(3128, "invalid price size")
	ErrNotOnlyMaker            = types.NewError(3129, "not only maker, kill order")
	ErrTradingUnavailable      = types.NewError(3130, "trading is unavailable in this market")
	ErrUserHasPendingTx        = types.NewError(3131, "user has pending transaction,such as order,stop,position")
	ErrPositionIsClosing       = types.NewError(3132, "position is closing")
	ErrPositionNumberExceed    = types.NewError(3133, "the number of positions exceeds the limit")
	ErrWithdrawableNotEnough   = types.NewError(3134, "withdrawable not enough")
)

func IsInvalidArguments(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpInvalidArgument) || errors.Is(err, ErrInvalidArgument)
}

func IsInternalError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpInternalError) || errors.Is(err, ErrInternalError)
}

func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpServiceUnavailable) || errors.Is(err, ErrServiceUnavailable)
}

func IsServiceTimeout(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpServiceTimeout) || errors.Is(err, ErrServiceTimeout)
}

func IsUnknownMethod(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpUnknownMethod) || errors.Is(err, ErrUnknownMethod)
}

func IsResultNull(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrHttpResultNull)
}

func IsRateLimit(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrRateLimit)
}
