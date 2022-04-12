package utils

type Int interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Float interface {
	float32 | float64
}

type Number interface {
	Int | Float
}

type Complex interface {
	complex64 | complex128
}

type AllNumber interface {
	Number | Complex
}
