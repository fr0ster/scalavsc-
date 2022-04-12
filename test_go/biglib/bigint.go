package biglib

import (
	"math/big"
)

type MyBigInt struct{ value *big.Int }

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *MyBigInt {
	return &MyBigInt{value: new(big.Int).SetInt64(x)}
}

func newFromInt(x *big.Int) *MyBigInt {
	return &MyBigInt{value: x}
}

func (v *MyBigInt) GetValue() IBigNumber {
	return v
}

func (v *MyBigInt) SetValue(x IBigNumber) IBigNumber {
	v = x.ToInt()
	return v
}

// For Interface IBigNumber
func (v *MyBigInt) ToFloat() *MyBigFloat {
	a := MyBigFloat{value: new(big.Float).SetInt(v.value)}
	return &a
}

func (v *MyBigInt) ToInt() *MyBigInt {
	return v
}

func (v *MyBigInt) FromFloat(x float64) IBigNumber {
	v.value = big.NewInt(int64(x))
	return v
}

func (v *MyBigInt) FromInt(x int64) IBigNumber {
	v.value = big.NewInt(x)
	return v
}

func (v *MyBigInt) Add(x IBigNumber) IBigNumber {
	a := v.ToInt().GetValue().ToInt().value
	b := x.ToInt().GetValue().ToInt().value
	return newFromInt(a.Add(a, b))
}

func (v *MyBigInt) Sub(x IBigNumber) IBigNumber {
	a := v.ToInt().GetValue().ToInt().value
	b := x.ToInt().GetValue().ToInt().value
	return newFromInt(a.Sub(a, b))
}

func (v *MyBigInt) Mul(x IBigNumber) IBigNumber {
	a := v.ToInt().GetValue().ToInt().value
	b := x.ToInt().GetValue().ToInt().value
	return newFromInt(a.Mul(a, b))
}

func (v *MyBigInt) Div(x IBigNumber) IBigNumber {
	a := v.ToInt().GetValue().ToInt().value
	b := x.ToInt().GetValue().ToInt().value
	return newFromInt(a.Mul(a, b))
}

func (v *MyBigInt) String() string {
	return v.value.String()
}

func (v *MyBigInt) compare(other *MyBigInt) int {
	return v.value.Cmp(other.value)
}
