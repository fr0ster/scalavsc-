package biglib

import "math/big"

type MyBigFloat struct{ value *big.Float }

// NewFloat allocates and returns a new Int set to x.
func NewFloat(x float64) *MyBigFloat {
	return &MyBigFloat{value: new(big.Float).SetFloat64(x)}
}

func (v *MyBigFloat) GetValue() *big.Float {
	return v.value
}

func (v *MyBigFloat) SetValue(NewValue *big.Float) {
	v.value = NewValue
}

// For Interface IBigNumber
func (v *MyBigFloat) ToFloat() *MyBigFloat {
	return v
}

func (v *MyBigFloat) ToInt() *MyBigInt {
	a := MyBigInt{value: v.ToInt().value}
	return &a
}

func (v *MyBigFloat) FromFloat(newValue *MyBigFloat) *MyBigFloat {
	v.value = newValue.value
	return v
}

func (v *MyBigFloat) FromInt(newValue *MyBigInt) *MyBigFloat {
	v = newValue.ToFloat()
	return v
}

func (v *MyBigFloat) String() string {
	return v.value.String()
}

func (v MyBigFloat) compare(other MyBigFloat) int {
	return v.value.Cmp(other.value)
}
