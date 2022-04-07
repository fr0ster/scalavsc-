package summator

import (
	"sync"
	. "test_go/utils"
)

func _summator(xs []int, ch chan int, wg *sync.WaitGroup) {
	summa := 0
	for _, i := range xs {
		summa += i
	}
	ch <- summa
	wg.Done()
}

func Summator(xs []int) int {
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
		go _summator(xs[i*step:(i+1)*step], ch[i], &wg)
	}
	wg.Wait()
	res := 0
	for _, v := range ch {
		res += <-v
		close(v)
	}
	return res
}
