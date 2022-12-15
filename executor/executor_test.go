package executor

import (
	"errors"
	"testing"
)

func fnSuccess(args ...interface{}) (interface{}, error) {
	return "string", nil
}

func fnFail(args ...interface{}) (interface{}, error) {
	return "", errors.New("this function fail")
}

func TestExec(t *testing.T) {
	executor := NewExecutor()
	executor.WithFunction(fnFail)
	executor.WithRetries(1)
	executor.Exec("test")
	err := executor.GetError()
	if err == nil {
		t.Fatal("executor should contains error")
	} else {
		t.Fatal(err)
	}
}
