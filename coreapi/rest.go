package coreapi

type Rest struct{}

type IRest interface{}

func NewRest() IRest {
	return Rest{}
}
