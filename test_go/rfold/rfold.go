package rfold

import (
	"sync"
)

func _go(n int, akk [][]int, xs []int) [][]int {
	if len(xs) <= n || n == 0 {
		return append(akk, xs)
	} else {
		return _go(n, append(akk, xs[:n]), xs[n:])
	}
}

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

func RFoldLeft(n int, xs []int, akk int, fn func(int, int) int) int {
	ch := make(chan int, n+1)
	var wg sync.WaitGroup
	xss := append([]int{akk}, xs...)
	if n >= len(xss) {
		n = 1
	}
	for _, line := range _go(int(len(xss)/n)+1, [][]int{}, xss) {
		wg.Add(1)
		go func(out chan<- int, xs []int, wg *sync.WaitGroup) {
			out <- _fold(xs[1:], xs[0], fn)
			wg.Done()
		}(ch, line, &wg)
	}
	wg.Wait()
	close(ch)
	res := <-ch
	wg.Add(1)
	go func() {
		for r := range ch {
			res = fn(res, r)
		}
		wg.Done()
	}()
	wg.Wait()
	return res
}
