package types

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type Number string

func (n Number) String() string {
	return string(n)
}

func (n Number) ToFloat() (float64, error) {
	return strconv.ParseFloat(n.String(), 64)
}

func (n Number) Float() float64 {
	v, _ := n.ToFloat()
	return v
}

func (n Number) ToDecimal() (decimal.Decimal, error) {
	return decimal.NewFromString(n.String())
}

func (n Number) Decimal() decimal.Decimal {
	v, _ := n.ToDecimal()
	return v
}

func (n Number) Truncate(precision int32) Number {
	return Number(n.Decimal().Truncate(precision).String())
}
