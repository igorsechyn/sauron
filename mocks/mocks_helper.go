package mocks

import "github.com/igorsechyn/sauron/pkg/app"

type Mocks struct {
	ConsoleWriter *MockConsoleWriter
	Metadata      *MockMetadata
	PluginService *MockPluginService
	FileSystem    *MockFileSystem
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
		ConsoleWriter: NewMockConsoleWriter(),
		Metadata:      NewMockMetadata(),
		PluginService: NewMockPluginService(),
		FileSystem:    NewMockFileSystem(),
	}
}
