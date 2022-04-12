package biglib

import (
	"math/big"
)

type myBigInt struct{ value *big.Int }

// NewInt allocates and returns a new Int set to x.
func newInt(x int64) IBigNumber {
	return &myBigInt{value: new(big.Int).SetInt64(x)}
}

func newFromInt(x *big.Int) IBigNumber {
	return &myBigInt{value: x}
}

func (v myBigInt) initValue() IBigNumber {
	a := myBigInt{value: big.NewInt(0)}
	return a
}

func (v myBigInt) GetValue() IBigNumber {
	return v
}

func (v myBigInt) SetValue(x IBigNumber) IBigNumber {
	v = x.toInt()
	return v
}

// For Interface IBigNumber
func (v myBigInt) toFloat() myBigFloat {
	a := myBigFloat{value: new(big.Float).SetInt(v.value)}
	return a
}

func (v myBigInt) toInt() myBigInt {
	return v
}

func (v myBigInt) fromFloat(x float64) IBigNumber {
	v.value = big.NewInt(int64(x))
	return v
}

func (v myBigInt) fromInt(x int64) IBigNumber {
	v.value = big.NewInt(x)
	return v
}

func (v myBigInt) Add(x IBigNumber) IBigNumber {
	a := v.GetValue().toInt().value
	b := x.GetValue().toInt().value
	return newFromInt(a.Add(a, b))
}

func (v myBigInt) Sub(x IBigNumber) IBigNumber {
	a := v.GetValue().toInt().value
	b := x.GetValue().toInt().value
	return newFromInt(a.Sub(a, b))
}

func (v myBigInt) Mul(x IBigNumber) IBigNumber {
	a := v.GetValue().toInt().value
	b := x.GetValue().toInt().value
	return newFromInt(a.Mul(a, b))
}

func (v myBigInt) Div(x IBigNumber) IBigNumber {
	a := v.GetValue().toInt().value
	b := x.GetValue().toInt().value
	return newFromInt(a.Mul(a, b))
}

func (v myBigInt) String() string {
	return v.value.String()
}

func (v myBigInt) Cmp(other IBigNumber) int {
	return v.value.Cmp(other.toInt().value)
}
