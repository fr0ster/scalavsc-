package biglib

type TBigNumber[T IBigNumber] struct {
	values T
}

// NewInt allocates and returns a new Int set to x.
func NewBigNumber[T IBigNumber](x T) *TBigNumber[T] {
	a := new(TBigNumber[T])
	a.values = x
	return a
}

func (v *TBigNumber[T]) GetInterface() *TBigNumber[T] {
	return v
}

func (v *TBigNumber[T]) GetValue() IBigNumber {
	return v.values.GetValue()
}

func (v *TBigNumber[T]) SetValue(NewValue IBigNumber) IBigNumber {
	return v.values.SetValue(NewValue)
}

// For Interface IBigNumber
func (v *TBigNumber[T]) ToFloat() *MyBigFloat {
	return v.values.ToFloat()
}

func (v *TBigNumber[T]) ToInt() *MyBigInt {
	return v.values.ToInt()
}

func (v *TBigNumber[T]) FromFloat(newValue float64) IBigNumber {
	return v.values.FromFloat(newValue)
}

func (v *TBigNumber[T]) FromInt(newValue int64) IBigNumber {
	return v.values.FromInt(newValue)
}

func (v *TBigNumber[T]) String() string {
	return v.values.String()
}

func (v *TBigNumber[T]) compare(other IBigNumber) int {
	return v.values.GetValue().ToFloat().compare(other.GetValue().ToFloat())
}
