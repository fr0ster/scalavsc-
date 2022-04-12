package biglib

import "math/big"

type MyBigFloat struct{ value *big.Float }

// NewFloat allocates and returns a new Int set to x.
func NewFloat(x float64) *MyBigFloat {
	return &MyBigFloat{value: new(big.Float).SetFloat64(x)}
}

func (v *MyBigFloat) GetValue() IBigNumber {
	return v
}

func (v *MyBigFloat) SetValue(x IBigNumber) IBigNumber {
	v = x.ToFloat()
	return v
}

// For Interface IBigNumber
func (v *MyBigFloat) ToFloat() *MyBigFloat {
	return v
}

func (v *MyBigFloat) ToInt() *MyBigInt {
	a := MyBigInt{value: v.ToInt().value}
	return &a
}

func (v *MyBigFloat) FromFloat(x float64) IBigNumber {
	v.value.SetFloat64(x)
	return v
}

func (v *MyBigFloat) FromInt(x int64) IBigNumber {
	v.value.SetInt64(x)
	return v
}

func (v *MyBigFloat) String() string {
	return v.value.String()
}

func (v *MyBigFloat) compare(other *MyBigFloat) int {
	return v.value.Cmp(other.value)
}
