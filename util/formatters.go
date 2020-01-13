package util

import (
	"math/big"
	"strconv"
)

func FormatTRBBalance(balance *big.Int) string {
	divisor := big.NewInt(1e18/100)
	divisor.Div(balance, divisor)
	f := float64(divisor.Uint64())/100
	return strconv.FormatFloat(f, 'f', 2, 64)
}
