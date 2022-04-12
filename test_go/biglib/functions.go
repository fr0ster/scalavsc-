package biglib

import (
	. "test_go/utils"
)

func Add[T Number](a T, b T) T {
	return a + b
}

// func Eq[T TBigNumber](a, b T) bool {
// 	// return a.ToFloat().Cmp(b.ToFloat()) == 0
// }

// func NeQ[T IBigInt](a, b T) bool {
// 	return a.ToFloat().value != b.ToFloat().value
// }

// func Test[T IBigInt](a, b T) bool {
// 	return a.ToFloat().compare(*b.ToFloat()) == 0
// }

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
