package util

import (
	"math/big"
	"strconv"
)

func FormatERC20Balance(balance *big.Int) string {
	f := 0.0
	if balance != nil {
		divisor := big.NewInt(1e18/100)
		divisor.Div(balance, divisor)
		f = float64(divisor.Uint64())/100
	}
	return strconv.FormatFloat(f, 'f', 2, 64)
}
