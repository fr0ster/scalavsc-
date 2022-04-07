package sum

import (
	"sync"
	. "test_go/utils"
)

func _sum(xs []int, akk *Akkum) {
	if len(xs) == 0 {
		akk.Wg.Done()
	} else {
		akk.Lock.Lock()
		akk.Sum += xs[0]
		akk.Lock.Unlock()
		_sum(xs[1:], akk)
	}
}

func Sum(xs []int) int {
	threads := THREADNUM
	step := len(xs)/threads + 1
	// limited size of arrays chunk
	if step > 25000000 {
		threads = THREADNUM * 1000
		step = len(xs)/threads + 1
	}
	ch := make([]chan int, threads)
	akk := Akkum{
		Sum:  0,
		Lock: sync.Mutex{},
		Wg:   sync.WaitGroup{},
	}
	for i := 0; i < threads; i++ {
		ch[i] = make(chan int, 1)
		akk.Wg.Add(1)
		end := (i + 1) * step
		if end > len(xs) {
			go _sum(xs[i*step:], &akk)
			break
		} else {
			go _sum(xs[i*step:(i+1)*step], &akk)
		}
	}
	akk.Wg.Wait()
	return akk.Sum
}
