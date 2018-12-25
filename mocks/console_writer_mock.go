package mocks

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockConsoleWriter struct {
	output string
	mock.Mock
}

func (writer *MockConsoleWriter) Write(bytes []byte) (int, error) {
	writer.Called(bytes)
	writer.output = fmt.Sprintf("%v%v", writer.output, string(bytes))
	return len(bytes), nil
}

func (writer *MockConsoleWriter) GetOutput() string {
	return writer.output
}

func NewMockConsoleWriter() *MockConsoleWriter {
	writer := new(MockConsoleWriter)
	writer.On("Write", mock.Anything).Return()
	return writer
}
