package mocks

import (
	"github.com/igorsechyn/sauron/pkg/app/metadata"
	"github.com/stretchr/testify/mock"
)

type MockMetadata struct {
	mock.Mock
}

func (mockMetadata *MockMetadata) Get() metadata.Info {
	args := mockMetadata.Called()
	return args.Get(0).(metadata.Info)
}

func (mockMetadata *MockMetadata) GivenGetReturns(info metadata.Info) {
	mockMetadata.ExpectedCalls = []*mock.Call{}
	mockMetadata.On("Get").Return(info)
}

func NewMockMetadata() *MockMetadata {
	mockMetadata := new(MockMetadata)
	mockMetadata.GivenGetReturns(metadata.Info{})
	return mockMetadata
}
