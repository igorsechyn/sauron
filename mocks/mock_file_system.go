package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockFileSystem struct {
	mock.Mock
}

func (mockFileSystem *MockFileSystem) Exists(path string) bool {
	args := mockFileSystem.Called(path)
	return args.Bool(0)
}

func (mockFileSystem *MockFileSystem) Save(path string, data io.Reader) error {
	args := mockFileSystem.Called(path, data)
	return args.Error(0)
}

func (mockFileSystem *MockFileSystem) GivenSaveSucceeds() {
	mockFileSystem.ExpectedCalls = getCallsWithoutMethod(mockFileSystem.ExpectedCalls, "Save")
	mockFileSystem.On("Save", mock.Anything, mock.Anything).Return(nil)
}

func (mockFileSystem *MockFileSystem) GivenSaveFailes(err error) {
	mockFileSystem.ExpectedCalls = getCallsWithoutMethod(mockFileSystem.ExpectedCalls, "Save")
	mockFileSystem.On("Save", mock.Anything, mock.Anything).Return(err)
}

func (mockFileSystem *MockFileSystem) GivenExistsReturns(exists bool) {
	mockFileSystem.ExpectedCalls = getCallsWithoutMethod(mockFileSystem.ExpectedCalls, "Exists")
	mockFileSystem.On("Exists", mock.Anything).Return(exists)
}

func NewMockFileSystem() *MockFileSystem {
	mockFileSystem := new(MockFileSystem)
	mockFileSystem.GivenExistsReturns(false)
	mockFileSystem.GivenSaveSucceeds()
	return mockFileSystem
}
