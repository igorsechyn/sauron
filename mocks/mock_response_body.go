package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockResponseBody struct {
	mock.Mock
	data []byte
}

func (mockResponseBody *MockResponseBody) Read(p []byte) (int, error) {
	n := copy(p, mockResponseBody.data)
	return n, nil
}

func (mockResponseBody *MockResponseBody) Close() error {
	mockResponseBody.Called()
	return nil
}

func NewMockResponseBody() *MockResponseBody {
	mockResponseBody := new(MockResponseBody)
	mockResponseBody.data = []byte("defalt-data")
	mockResponseBody.On("Close").Return()
	return mockResponseBody
}
