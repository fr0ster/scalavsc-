package biglib

type GBigNumber interface {
	MyBigInt | MyBigFloat
}

type IBigNumber interface {
	ToFloat() *MyBigFloat
	ToInt() *MyBigInt
	GetValue() IBigNumber
	SetValue(IBigNumber) IBigNumber
	FromFloat(newValue float64) IBigNumber
	FromInt(newValue int64) IBigNumber
	Add(IBigNumber) IBigNumber
	Sub(IBigNumber) IBigNumber
	Mul(IBigNumber) IBigNumber
	Div(IBigNumber) IBigNumber
	String() string
}
