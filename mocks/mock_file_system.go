package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockFileSystem struct {
	mock.Mock
}

func (mockFileSystem *MockFileSystem) Exists(path string) bool {
	args := mockFileSystem.Called(path)
	return args.Bool(0)
}

func (mockFileSystem *MockFileSystem) GivenExistsReturns(exists bool) {
	mockFileSystem.ExpectedCalls = []*mock.Call{}
	mockFileSystem.On("Exists", mock.Anything).Return(exists)
}

func NewMockFileSystem() *MockFileSystem {
	mockFileSystem := new(MockFileSystem)
	mockFileSystem.GivenExistsReturns(false)
	return mockFileSystem
}
