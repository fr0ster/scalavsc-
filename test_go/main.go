package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

func _factor(akk *big.Int, n int64) *big.Int {
	if n == 1 {
		return akk
	} else {
		return _factor(big.NewInt(0).Mul(akk, big.NewInt(n)), n-1)
	}
}

func Fac(n int64) *big.Int {
	return _factor(big.NewInt(1), n)
}

func _summator(id int, wg *sync.WaitGroup, ch chan<- int, akk int, xs []int) {
	if len(xs) == 0 {
		ch <- akk
	} else {
		if len(xs)%10 == 0 {
			_summator(id+1, wg, ch, akk+xs[0], xs[1:])
		} else {
			go func() {
				wg.Add(1)
				_summator(id+1, wg, ch, akk+xs[0], xs[1:])
			}()
		}
	}
}

func Sum(xs []int) int {
	ch := make(chan int)
	var wg sync.WaitGroup
	go _summator(0, &wg, ch, 0, xs)
	wg.Wait()
	return <-ch
}

func Summator(xs []int) int {
	summa := 0
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		for _, i := range xs {
			summa += i
		}
		ch <- summa
		wg.Done()
	}()
	wg.Wait()
	res := <-ch
	close(ch)
	return res
}

func _go(n int, akk [][]int, xs []int) [][]int {
	if len(xs) <= n || n == 0 {
		return append(akk, xs)
	} else {
		return _go(n, append(akk, xs[:n]), xs[n:])
	}
}

func _fold(xs []int, akk int, fn func(a int, b int) int) int {
	if len(xs) == 0 {
		return akk
	} else {
		if len(xs) == 1 {
			return fn(akk, xs[0])
		} else {
			return _fold(xs[1:], fn(xs[0], akk), fn)
		}
	}
}

func FoldLeft(xs []int, akk int, fn func(a int, b int) int) int {
	return _fold(xs, akk, fn)
}

func RFoldLeft(n int, xs []int, akk int, fn func(int, int) int) int {
	ch := make(chan int, n+1)
	var wg sync.WaitGroup
	xss := append([]int{akk}, xs...)
	if n >= len(xss) {
		n = 1
	}
	for _, line := range _go(int(len(xss)/n)+1, [][]int{}, xss) {
		wg.Add(1)
		go func(out chan<- int, xs []int, wg *sync.WaitGroup) {
			out <- _fold(xs[1:], xs[0], fn)
			wg.Done()
		}(ch, line, &wg)
	}
	wg.Wait()
	close(ch)
	res := <-ch
	wg.Add(1)
	go func() {
		for r := range ch {
			res = fn(res, r)
		}
		wg.Done()
	}()
	wg.Wait()
	return res
}

const GOMAX = 2

func main() {
	fmt.Println("vim-go")
	_gen := func(n int) []int {
		res := []int{}
		i := 1
		for i <= n {
			res = append(res, i)
			i += 1
		}
		return res
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "-g" {
			a := []int{}
			for _, v := range os.Args[2:] {
				if b, err := strconv.Atoi(v); err == nil {
					a = append(a, b)
				}
			}
			fmt.Printf("Summator(%v) = %v\n", a, Summator(a))
			fmt.Printf("Sum(%v) = %v\n", a, Sum(a))
			fmt.Printf("_go(%v) = %v\n", a, _go(int(len(a)/5), [][]int{}, a))
			fmt.Printf("_go(%v) = %v\n", a, _go(5, [][]int{}, a))
			fmt.Printf("Go - FoldLeft(%v,{a+b}) = %v\n", a, FoldLeft(a[1:], a[0], func(a int, b int) int { return a + b }))
			fmt.Printf("Go - RFoldLeft(%v,{a+b}) = %v\n", a, RFoldLeft(5, a[1:], a[0], func(a int, b int) int { return a + b }))
			fmt.Printf("Go - Sum(%v) = %v\n", a, FoldLeft(a[1:], a[0], func(a int, b int) int { return a + b }))
			fmt.Printf("Go - Sum(%v) = %v\n", a, RFoldLeft(GOMAX, a[1:], a[0], func(a int, b int) int { return a + b }))
		}
		if os.Args[1] == "-n" {
			if v, err := strconv.Atoi(os.Args[2]); err == nil {
				xs := _gen(v)
				fmt.Printf("Go - Sum = %v\n", RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Mul = %v\n", RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a * b }))
			}
		}
		if os.Args[1] == "-t" {
			toval := int(2147483647 / 10)
			if v, err := strconv.Atoi(os.Args[2]); err == nil {
				//fmt.Printf("Go - Sum = %v\n", RFoldLeft(v, _gen(5000000)[1:], 1, func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Sum = %v\n", RFoldLeft(v, _gen(toval)[1:], 1, func(a int, b int) int { return a + b }))
			}
		}
	} else {
		fmt.Printf("Max Int == {}\n", int((^uint(0))>>1))
		fmt.Print("\n")
		fmt.Printf("Go - Sum = %v\n", RFoldLeft(GOMAX, _gen(5000000)[1:], 1, func(a int, b int) int { return a + b }))
	}
}
