package mocks

import (
	"github.com/igorsechyn/sauron/pkg/app/plugin"
	"github.com/stretchr/testify/mock"
)

type MockPluginInstaller struct {
	mock.Mock
}

func (mockInstaller *MockPluginInstaller) Install(info plugin.Info) error {
	args := mockInstaller.Called(info)
	return args.Error(0)
}

func (mockInstaller *MockPluginInstaller) GivenInstallSucceeds() {
	mockInstaller.On("Install", mock.AnythingOfType("plugin.Info")).Return(nil)
}

func (mockInstaller *MockPluginInstaller) GivenInstallFails(err error) {
	mockInstaller.ExpectedCalls = []*mock.Call{}
	mockInstaller.On("Install", mock.AnythingOfType("plugin.Info")).Return(err)
}

func NewMockPluginInstaller() *MockPluginInstaller {
	mockInstaller := new(MockPluginInstaller)
	mockInstaller.GivenInstallSucceeds()
	return mockInstaller
}
