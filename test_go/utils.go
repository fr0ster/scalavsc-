package main

import (
	"math/rand"
	"time"
)

func _gen(n int) []int {
	res := []int{}
	rand.Seed(time.Now().UnixNano())
	i := 1
	for i <= n {
		res = append(res, rand.Intn(n))
		i++
	}
	return res
}

func _split(xs []int, step int) [][]int {
	res := [][]int{}
	for i := 0; i < len(xs); i += step {
		end := (i + step)
		if end > len(xs) {
			res = append(res, xs[i:])
		} else {
			res = append(res, xs[i:end])
		}
	}
	return res
}
