package executor

import (
	"errors"
	"fmt"
)

var (
	ReachLimitErr = errors.New("retries execution reach limit")
)

type Executor struct {
	retries int
	fn      func(args ...interface{}) (interface{}, error)
	err     error
	result  interface{}
}

type IExector interface {
	WithRetries(retries int)
	WithFunction(fn func(args ...interface{}) (interface{}, error))
	Exec(args ...interface{}) (interface{}, error)
	GetResult() interface{}
	GetError() error
}

func NewExecutor() IExector {
	return &Executor{}
}

func (e *Executor) WithRetries(retries int) {
	e.retries = retries
}

func (e *Executor) WithFunction(fn func(args ...interface{}) (interface{}, error)) {
	e.fn = fn
}

func (e *Executor) Reset() {
	e.retries = 0
	e.fn = nil
	e.err = nil
	e.result = nil
}

func (e *Executor) Success(result interface{}) {
	e.Reset()
	e.result = result
}

func (e *Executor) Fail(err error) {
	e.Reset()
	e.err = err
}

func (e *Executor) Exec(args ...interface{}) (interface{}, error) {
	result, err := e.fn(args)
	cur := 0
	for cur < e.retries && err != nil {
		cur++
		result, err = e.fn(args)
	}
	if err != nil {
		errorWrap := fmt.Errorf("%w: %d, %s", ReachLimitErr, cur, err.Error())
		e.Fail(errorWrap)
		return result, errorWrap
	}
	e.Success(result)
	return result, nil
}

func (e *Executor) GetResult() interface{} {
	return e.result
}

func (e *Executor) GetError() error {
	return e.err
}
