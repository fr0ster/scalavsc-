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
	if len(xs) > THREADNUM*10 {
		step = len(xs) / THREADNUM
	}
	ch := []chan int{}
	_xs := Split(xs, step)
	var wg sync.WaitGroup
	for _, v := range _xs {
		_ch := make(chan int, 1)
		ch = append(ch, _ch)
		wg.Add(1)
		v := v
		go _summator(v, _ch, &wg)
	}
	wg.Wait()
	res := 0
	for _, v := range ch {
		res += <-v
		close(v)
	}
	return res
}
