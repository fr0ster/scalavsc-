package main

import (
	"sync"
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
	if len(xs) > THREADNUM*10 {
		step = len(xs) / THREADNUM
	}
	ch_arr := []chan int{}
	var wg sync.WaitGroup
	_xs := _split(xs, step)
	for _, part := range _xs {
		xs = part
		ch := make(chan int, 1)
		ch_arr = append(ch_arr, ch)
		wg.Add(1)
		go _sum(0, part, ch, &wg)
	}
	wg.Wait()
	res := 0
	for _, _ch := range ch_arr {
		res += <-_ch
	}
	return res
}
