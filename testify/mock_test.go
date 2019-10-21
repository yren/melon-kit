package testify

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

// MyMockedObject is mock object
type MyMockedObject struct{
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestSome(t *testing.T) {
	testObj := new(MyMockedObject)
	testObj.On("DoSomething", 123).Return(true, nil)
}