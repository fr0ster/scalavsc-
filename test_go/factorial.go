package main

import (
	"math/big"
)

func _factor(akk *big.Int, n int64) *big.Int {
	if n == 1 {
		return akk
	} else {
		return _factor(big.NewInt(0).Mul(akk, big.NewInt(n)), n-1)
	}
}

func Fac(n int64) *big.Int {
	return _factor(big.NewInt(1), n)
}
