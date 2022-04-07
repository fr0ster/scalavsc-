package utils

import (
	"math/rand"
	"time"
)

const (
	GOMAX     = 2
	THREADNUM = 4
	STEPS     = 1000
)

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
