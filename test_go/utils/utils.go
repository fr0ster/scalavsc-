package utils

import (
	"math/rand"
	"sync"
	"time"
)

type Akkum struct {
	Sum  int
	Lock sync.Mutex
	Wg   sync.WaitGroup
}

const (
	GOMAX     = 2
	THREADNUM = 4
	STEPS     = 10000
)

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128
}

type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 |
		~string
}

func Gen(n int) []int {
	res := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(res); i++ {
		res[i] = rand.Intn(n)
	}
	return res
}

func Split(xs []int, step int) [][]int {
	n := len(xs) / step
	var res [][]int
	for i := 0; i < n; i++ {
		res = append(res, make([]int, step))
		end := (i + 1) * step
		if end >= len(xs) {
			res[i] = xs[i*step:]
		} else {
			res[i] = xs[i*step : end]
		}
	}
	return res
}
