package biglib

type GBigNumber interface {
	myBigInt | myBigFloat
}

type IBigNumber interface {
	toFloat() myBigFloat
	toInt() myBigInt
	fromFloat(newValue float64) IBigNumber
	fromInt(newValue int64) IBigNumber
	initValue() IBigNumber
	GetValue() IBigNumber
	SetValue(IBigNumber) IBigNumber
	Add(IBigNumber) IBigNumber
	Sub(IBigNumber) IBigNumber
	Mul(IBigNumber) IBigNumber
	Div(IBigNumber) IBigNumber
	String() string
	Cmp(IBigNumber) int
}
