package sum

import (
	"sync"
	. "test_go/utils"
)

func _sum(akk int, xs []int, ch chan int, wg *sync.WaitGroup) {
	if len(xs) == 0 {
		ch <- akk
		wg.Done()
	} else {
		_sum(akk+xs[0], xs[1:], ch, wg)
	}
}

func Sum(xs []int) int {
	step := len(xs)
	threads := 1
	if len(xs) > THREADNUM*10 {
		step = len(xs) / THREADNUM
		threads = THREADNUM
	}
	ch := make([]chan int, threads)
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		ch[i] = make(chan int, 1)
		wg.Add(1)
		go _sum(0, xs[i*step:(i+1)*step], ch[i], &wg)
	}
	wg.Wait()
	res := 0
	for _, _ch := range ch {
		res += <-_ch
	}
	return res
}
