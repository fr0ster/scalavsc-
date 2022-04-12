package biglib

type IBigNumber interface {
	ToFloat() *MyBigFloat
	ToInt() *MyBigInt
	GetValue() IBigNumber
	SetValue(IBigNumber) IBigNumber
	FromFloat(newValue float64) IBigNumber
	FromInt(newValue int64) IBigNumber
	String() string
}
