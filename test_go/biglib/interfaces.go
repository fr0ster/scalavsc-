package biglib

type BigNumberModel interface {
	ID() string
}

type BigNumberProvider[MODEL BigNumberModel] interface {
	FindByID(id string) (MODEL, error)
	List() ([]MODEL, error)
	Update(id string, model MODEL) error
	Insert(model MODEL) error
	Delete(id string) error
	String() string
	// 	// comparable
}

type IBigNumber[T GBigNumber] interface {
	ToFloat() *MyBigFloat
	ToInt() *MyBigInt
	GetValue() *T
	SetValue(*T) *TBigNumber[T]
	FromFloat(newValue float64) *T
	FromInt(newValue int64) *T
	String() string
	// comparable
}
