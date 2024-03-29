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
	threads := THREADNUM
	step := len(xs)/threads + 1
	ch := make([]chan int, threads)
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		ch[i] = make(chan int, 1)
		wg.Add(1)
		end := (i + 1) * step
		if end > len(xs) {
			go _summator(xs[i*step:], ch[i], &wg)
			break
		} else {
			go _summator(xs[i*step:(i+1)*step], ch[i], &wg)
		}
	}
	wg.Wait()
	res := 0
	for _, _ch := range ch {
		if _ch != nil {
			res += <-_ch
			close(_ch)
		}
	}
	return res
}
