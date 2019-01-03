package mocks

import (
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/stretchr/testify/mock"
)

type Mocks struct {
	ConsoleWriter *MockConsoleWriter
	Metadata      *MockMetadata
	PluginService *MockPluginService
	FileSystem    *MockFileSystem
	HttpClient    *MockHttpClient
}

func (allMocks *Mocks) ToAppBoundaries() app.Boundaries {
	return app.Boundaries{
		ConsoleWriter: allMocks.ConsoleWriter,
		Metadata:      allMocks.Metadata,
		FileSystem:    allMocks.FileSystem,
		HttpClient:    allMocks.HttpClient,
	}
}

func InitAllMocks() *Mocks {
	return &Mocks{
		ConsoleWriter: NewMockConsoleWriter(),
		Metadata:      NewMockMetadata(),
		PluginService: NewMockPluginService(),
		FileSystem:    NewMockFileSystem(),
		HttpClient:    NewMockHttpClient(),
	}
}

func getCallsWithoutMethod(calls []*mock.Call, method string) []*mock.Call {
	callsWithoutMethod := make([]*mock.Call, 0, 0)
	for _, call := range calls {
		if call.Method != method {
			callsWithoutMethod = append(callsWithoutMethod, call)
		}
	}

	return callsWithoutMethod
}
