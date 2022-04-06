package fold

func _fold(xs []int, akk int, fn func(a int, b int) int) int {
	if len(xs) == 0 {
		return akk
	} else {
		return _fold(xs[1:], fn(xs[0], akk), fn)
	}
}

func FoldLeft(xs []int, akk int, fn func(a int, b int) int) int {
	return _fold(xs, akk, fn)
}
