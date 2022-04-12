package biglib

import (
	"math"
	. "test_go/utils"
	"testing"
)

func gTest1[T Float](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

func gTest2[T Int](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

func gTest3[T Number](a T, b T) T {
	la := float64(a)
	lb := float64(b)
	return T(math.Pow(la, lb))
}

// func gTest4[T IBigInt](a T, b T) *MyBigFloat {
// 	la := a.ToFloat().value
// 	lb := b.ToFloat().value
// 	return &MyBigFloat{value: la.Add(la, lb)}
// }

func TestNewT(t *testing.T) {
	// t.Log("gTest1[float32](10.0, 2.0)", gTest1[float32](10.0, 2.0))
	// t.Error("gTest1[float32](10.0, 2.0)", gTest1[float32](10.0, 2.0))
	// t.Log("gTest2[int32](10.0, 2.0)", gTest2[int32](10.0, 2.0))
	// t.Error("gTest2[int32](10.0, 2.0)", gTest2[int32](10.0, 2.0))
	// t.Log("gTest3[float32](10.0, 2.0)", gTest3[float32](10.0, 2.0))
	// t.Error("gTest3[float32](10.0, 2.0)", gTest3[float32](10.0, 2.0))
	// a := NewInt(1000)
	// t.Error(a)
	// a.FromInt(NewInt(2000))
	// t.Error(a)
	// t.Error(gTest4(NewInt(1000), NewInt(1000)))
	// t.Error(Test(NewInt(1000), NewInt(1000)))
	a := NewBigNumber(NewInt(10000)).values
	t.Error(a)
	b := NewBigNumber(NewFloat(10000))
	t.Error(b.values)
	c := b.SetValue(NewFloat(20000)).GetValue()
	t.Error(b.values)
	t.Error(c)
}
