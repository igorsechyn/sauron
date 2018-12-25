package mocks

type MockBoundaries struct {
	ConsoleWriter   *MockConsoleWriter
	Metadata        *MockMetadata
	PluginInstaller *MockPluginInstaller
}

func InitAllMocks() *MockBoundaries {
	return &MockBoundaries{
		ConsoleWriter:   NewMockConsoleWriter(),
		Metadata:        NewMockMetadata(),
		PluginInstaller: NewMockPluginInstaller(),
	}
}
