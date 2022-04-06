package main

// func _go(n int, akk [][]int, xs []int) [][]int {
// 	if len(xs) <= n || n == 0 {
// 		return append(akk, xs)
// 	} else {
// 		return _go(n, append(akk, xs[:n]), xs[n:])
// 	}
// }

func _fold(xs []int, akk int, fn func(a int, b int) int) int {
	if len(xs) == 0 {
		return akk
	} else {
		if len(xs) == 1 {
			return fn(akk, xs[0])
		} else {
			return _fold(xs[1:], fn(xs[0], akk), fn)
		}
	}
}

func FoldLeft(xs []int, akk int, fn func(a int, b int) int) int {
	return _fold(xs, akk, fn)
}
