package biglib

import "math/big"

type myBigFloat struct{ value *big.Float }

// NewFloat allocates and returns a new Int set to x.
func newFloat(x float64) IBigNumber {
	return &myBigFloat{value: new(big.Float).SetFloat64(x)}
}

func newFromFloat(x *big.Float) IBigNumber {
	return &myBigFloat{value: x}
}

func (v myBigFloat) initValue() IBigNumber {
	a := myBigFloat{value: big.NewFloat(0)}
	return a
}

func (v myBigFloat) GetValue() IBigNumber {
	return v
}

func (v myBigFloat) SetValue(x IBigNumber) IBigNumber {
	v = x.toFloat()
	return v
}

// For Interface IBigNumber
func (v myBigFloat) toFloat() myBigFloat {
	return v
}

func (v myBigFloat) toInt() myBigInt {
	a := myBigInt{value: v.toInt().value}
	return a
}

func (v myBigFloat) fromFloat(x float64) IBigNumber {
	(v.value).SetFloat64(x)
	return v
}

func (v myBigFloat) fromInt(x int64) IBigNumber {
	v.value.SetInt64(x)
	return v
}

func (v myBigFloat) Add(x IBigNumber) IBigNumber {
	a := v.GetValue().toFloat().value
	b := x.GetValue().toFloat().value
	return newFromFloat(a.Add(a, b))
}

func (v myBigFloat) Sub(x IBigNumber) IBigNumber {
	a := v.GetValue().toFloat().value
	b := x.GetValue().toFloat().value
	return newFromFloat(a.Sub(a, b))
}

func (v myBigFloat) Mul(x IBigNumber) IBigNumber {
	a := v.GetValue().toFloat().value
	b := x.GetValue().toFloat().value
	return newFromFloat(a.Mul(a, b))
}

func (v myBigFloat) Div(x IBigNumber) IBigNumber {
	a := v.GetValue().toFloat().value
	b := x.GetValue().toFloat().value
	return newFromFloat(a.Mul(a, b))
}

func (v myBigFloat) String() string {
	return v.value.String()
}

func (v myBigFloat) Cmp(other IBigNumber) int {
	return v.value.Cmp(other.toFloat().value)
}
