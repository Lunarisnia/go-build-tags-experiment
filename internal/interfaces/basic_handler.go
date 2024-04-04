package interfaces

type BasicHandler[ST any] interface {
	GetRequestStruct() (ST, error)
}
