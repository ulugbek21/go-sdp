package future

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// SuccessFunc ...
type SuccessFunc func(string)

// FailFunc ...
type FailFunc func(error)

// ExecuteStringFunc ...
type ExecuteStringFunc func() (string, error)

// MaybeString ...
type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

// Success ...
func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

// Fail ...
func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

// Execute ...
func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}
