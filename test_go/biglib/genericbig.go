package biglib

import "math/big"

type TBigNumber[T IBigNumber] struct {
	values T
}

// NewInt allocates and returns a new Int set to x.
func NewBigNumber[T IBigNumber]() IBigNumber {
	a := *new(T)
	return a.initValue()
}

func NewFromInt(x int64) IBigNumber {
	return newInt(x)
}

func NewFromFloat(x float64) IBigNumber {
	return newFloat(x)
}

func NewFromBigInt(x *big.Int) IBigNumber {
	return newFromInt(x)
}

func NewFromBigFloat(x *big.Float) IBigNumber {
	return newFromFloat(x)
}

func (v *TBigNumber[T]) initValue() IBigNumber {
	a := *new(T)
	return a.initValue()
}

func (v *TBigNumber[T]) GetValue() IBigNumber {
	return v.values.GetValue()
}

func (v *TBigNumber[T]) SetValue(NewValue IBigNumber) IBigNumber {
	return v.values.SetValue(NewValue)
}

// For Interface IBigNumber
func (v *TBigNumber[T]) toFloat() myBigFloat {
	return v.values.toFloat()
}

func (v *TBigNumber[T]) toInt() myBigInt {
	return v.values.toInt()
}

func (v *TBigNumber[T]) FromFloat(newValue float64) IBigNumber {
	return v.values.fromFloat(newValue)
}

func (v *TBigNumber[T]) FromInt(newValue int64) IBigNumber {
	return v.values.fromInt(newValue)
}

func (v *TBigNumber[T]) String() string {
	return v.values.String()
}

func (v *TBigNumber[T]) Cmp(other IBigNumber) int {
	return v.values.GetValue().toFloat().Cmp(other.GetValue().toFloat())
}
