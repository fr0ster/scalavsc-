package biglib

type TBigNumber[T GBigNumber] struct {
	values *T
}

// NewInt allocates and returns a new Int set to x.
func NewBigNumber[T GBigNumber](x *T) *TBigNumber[T] {
	a := new(TBigNumber[T])
	a.values = x
	return a
}

func (v *TBigNumber[T]) GetValue() *T {
	return v.values
}

func (v *TBigNumber[T]) SetValue(NewValue *T) *TBigNumber[T] {
	v.values = NewValue
	return v
}

// // For Interface IBigNumber
// func (v *TBigNumber[T]) ToFloat() *MyBigFloat {
// 	a := MyBigFloat{value: v.values}
// 	return v.values.ToFloat()
// }

// func (v *MyBigInt) ToInt() *MyBigInt {
// 	return v
// }

// func (v *TBigNumber[GBigNumber, IBigNumber]) FromFloat(newValue float64) *TBigNumber[GBigNumber, IBigNumber] {
// 	// a := NewBigNumber[GBigNumber, IBigNumber](x)
// 	return v
// }

// func (v *MyBigInt) FromInt(newValue *MyBigInt) *MyBigInt {
// 	v.value = newValue.value
// 	return v
// }

// func (v *TBigNumber[GBigNumber]) String() string {
// 	return v.GetValue().GetValue().String()
// }

// func (v MyBigInt) compare(other MyBigInt) int {
// 	return v.value.Cmp(other.value)
// }
