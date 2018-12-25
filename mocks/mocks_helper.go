package mocks

import "github.com/igorsechyn/sauron/pkg/app"

type Mocks struct {
	ConsoleWriter   *MockConsoleWriter
	Metadata        *MockMetadata
	PluginInstaller *MockPluginInstaller
	FileSystem      *MockFileSystem
}

func (allMocks *Mocks) ToAppBoundaries() app.Boundaries {
	return app.Boundaries{
		ConsoleWriter: allMocks.ConsoleWriter,
		Metadata:      allMocks.Metadata,
		FileSystem:    allMocks.FileSystem,
	}
}

func InitAllMocks() *Mocks {
	return &Mocks{
		ConsoleWriter:   NewMockConsoleWriter(),
		Metadata:        NewMockMetadata(),
		PluginInstaller: NewMockPluginInstaller(),
		FileSystem:      NewMockFileSystem(),
	}
}
