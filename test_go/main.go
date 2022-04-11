package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	. "test_go/bigfib"
	. "test_go/fib"
	. "test_go/fold"
	. "test_go/rfold"
	. "test_go/sc/lock/sum"
	. "test_go/sc/lock/summator"
	. "test_go/utils"
)

func main() {
	fmt.Println("vim-go")
	if len(os.Args) > 1 {
		if os.Args[1] == "-g" {
			a := []int{}
			for _, v := range os.Args[2:] {
				if b, err := strconv.Atoi(v); err == nil {
					a = append(a, b)
				}
			}
			if len(a) > 0 {
				fmt.Printf("Summator(%v) = %v\n", a, Summator(a))
				fmt.Printf("Sum(%v) = %v\n", a, Sum(a))
				fmt.Printf("Go - FoldLeft(%v,{a+b}) = %v\n", a, FoldLeft(a[1:], a[0], func(a int, b int) int { return a + b }))
				fmt.Printf("Go - RFoldLeft(%v,{a+b}) = %v\n", a, RFoldLeft(5, a[1:], a[0], func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Sum(%v) = %v\n", a, FoldLeft(a[1:], a[0], func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Sum(%v) = %v\n", a, RFoldLeft(GOMAX, a[1:], a[0], func(a int, b int) int { return a + b }))
			}
		}
		if os.Args[1] == "-n" {
			if v, err := strconv.Atoi(os.Args[2]); err == nil {
				xs := Gen(v)
				fmt.Printf("Go - Sum = %v\n", RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Mul = %v\n", RFoldLeft(GOMAX, xs[1:], xs[0], func(a int, b int) int { return a * b }))
			}
		}
		if os.Args[1] == "-t" {
			if v, err := strconv.Atoi(os.Args[2]); err == nil {
				toval := int(2147483647 / 10)
				debug.SetGCPercent(-1)
				//fmt.Printf("Go - Sum = %v\n", RFoldLeft(v, Gen(5000000)[1:], 1, func(a int, b int) int { return a + b }))
				fmt.Printf("Go - Sum = %v\n", RFoldLeft(v, Gen(toval)[1:], 1, func(a int, b int) int { return a + b }))
			}
		}
		if os.Args[1] == "-l" && len(os.Args) > 2 {
			if v, err := strconv.Atoi(os.Args[2]); err == nil {
				xs := Gen(1000)
				for i := 0; i < v; i++ {
					fmt.Printf("Go - Sum %v = %v\n", v, Sum(xs))
				}
			}
		}
		if os.Args[1] == "-s" {
			xs := Split(Gen(9), 5)
			fmt.Println(xs)
		}
		if os.Args[1] == "-fib" {
			n := int64(93)
			fmt.Println("BigFib(", n, ") == ", BigFib(n))
			fmt.Println(BigFib(5).ProbablyPrime(20))
			fmt.Println("Fib(", n, ") == ", Fib[int64, uint64](n))
			c := uintptr(n)
			fmt.Println("Fib(", n, ") == ", Fib[uintptr, uint64](c))
		}
	} else {
		fmt.Println("Max Int == {}\n", int((^uint(0))>>1))
		fmt.Print("\n")
		fmt.Printf("Go - Sum = %v\n", RFoldLeft(GOMAX, Gen(5000000)[1:], 1, func(a int, b int) int { return a + b }))
	}
}
