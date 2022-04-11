package fib

import "math/big"

// A tail recursive function to
// calculate n th fibonacci number
func _bigfib(n int64, a *big.Int, b *big.Int) *big.Int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	c := a.Add(a, b)
	return _bigfib(n-1, b, c)
}
func BigFib(n int64) *big.Int {
	return _bigfib(n, big.NewInt(0), big.NewInt(1))
}
