package mocks

import (
	"github.com/igorsechyn/sauron/pkg/app/plugin"
	"github.com/stretchr/testify/mock"
)

type MockPluginService struct {
	mock.Mock
}

func (mockService *MockPluginService) Install(info plugin.Info) error {
	args := mockService.Called(info)
	return args.Error(0)
}

func (mockService *MockPluginService) GivenInstallSucceeds() {
	mockService.ExpectedCalls = getCallsWithoutMethod(mockService.ExpectedCalls, "Install")
	mockService.On("Install", mock.AnythingOfType("plugin.Info")).Return(nil)
}

func (mockService *MockPluginService) GivenInstallFails(err error) {
	mockService.ExpectedCalls = getCallsWithoutMethod(mockService.ExpectedCalls, "Install")
	mockService.On("Install", mock.AnythingOfType("plugin.Info")).Return(err)
}

func NewMockPluginService() *MockPluginService {
	mockService := new(MockPluginService)
	mockService.GivenInstallSucceeds()
	return mockService
}
