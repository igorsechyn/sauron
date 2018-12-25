package mocks

import (
	"github.com/igorsechyn/sauron/pkg/app/metadata"
	"github.com/stretchr/testify/mock"
)

type MockMetadata struct {
	mock.Mock
}

func (mock *MockMetadata) Get() metadata.Info {
	args := mock.Called()
	return args.Get(0).(metadata.Info)
}

func (mock *MockMetadata) GivenGetReturns(info metadata.Info) {
	mock.On("Get").Return(info)
}

func NewMockMetadata() *MockMetadata {
	mockMetadata := new(MockMetadata)
	mockMetadata.GivenGetReturns(metadata.Info{})
	return mockMetadata
}
