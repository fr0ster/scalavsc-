package biglib

import (
	"math/big"
)

type MyBigInt struct{ value *big.Int }

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *MyBigInt {
	return &MyBigInt{value: new(big.Int).SetInt64(x)}
}

func (v *MyBigInt) GetValue() *big.Int {
	return v.value
}

func (v *MyBigInt) SetValue(NewValue *big.Int) {
	v.value = NewValue
}

// For Interface IBigNumber
func (v *MyBigInt) ToFloat() *MyBigFloat {
	a := MyBigFloat{value: new(big.Float).SetInt(v.value)}
	return &a
}

func (v *MyBigInt) ToInt() *MyBigInt {
	return v
}

func (v *MyBigInt) FromFloat(newValue *MyBigFloat) *MyBigInt {
	v.value, _ = new(big.Int).SetString(newValue.value.Text('f', 0), 10)
	return v
}

func (v *MyBigInt) FromInt(newValue *MyBigInt) *MyBigInt {
	v.value = newValue.value
	return v
}

func (v *MyBigInt) String() string {
	return v.value.String()
}

func (v MyBigInt) compare(other MyBigInt) int {
	return v.value.Cmp(other.value)
}
