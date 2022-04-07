package summator

import (
	"sync"
	. "test_go/utils"
)

func _summator(xs []int, akk *Akkum) {
	res := 0
	for i := 0; i < len(xs); i++ {
		res += xs[i]
	}
	akk.Lock.Lock()
	akk.Sum += res
	akk.Lock.Unlock()
	akk.Wg.Done()
}

func Summator(xs []int) int {
	step := len(xs) / THREADNUM
	akk := Akkum{
		Sum:  0,
		Lock: sync.Mutex{},
		Wg:   sync.WaitGroup{},
	}
	for i := 0; i < len(xs); i += step {
		end := i + step
		akk.Wg.Add(1)
		if end >= len(xs) {
			go _summator(xs[i:], &akk)
			break
		} else {
			go _summator(xs[i:i+step], &akk)
		}
	}
	akk.Wg.Wait()
	return akk.Sum
}
