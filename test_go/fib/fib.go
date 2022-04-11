package fib

import "test_go/utils"

// A tail recursive function to
// calculate n th fibonacci number
func _fib[A, B utils.Int](n A, a B, b B) B {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	return _fib(n-1, b, a+b)
}
func Fib[A, B utils.Int](n A) B {
	return _fib[A, B](n, 0, 1)
}
