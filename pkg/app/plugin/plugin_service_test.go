package plugin_test

import (
	"testing"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/app/metadata"
	"github.com/igorsechyn/sauron/pkg/app/plugin"
)

type installOption func(pluginInfo *plugin.Info)

func pluginName(name string) installOption {
	return func(pluginInfo *plugin.Info) {
		pluginInfo.PluginName = name
	}
}

func version(version string) installOption {
	return func(pluginInfo *plugin.Info) {
		pluginInfo.Version = version
	}
}

func url(url string) installOption {
	return func(pluginInfo *plugin.Info) {
		pluginInfo.URL = url
	}
}

func whenInstallPluginIsInvoked(allMocks *mocks.Mocks, options ...installOption) {
	pluginInfo := plugin.Info{
		PluginName:       "default-plugin-name",
		LongDescription:  "default-long-description",
		ShortDescription: "default-short-description",
		URL:              "http://default.url",
		Version:          "default",
	}

	for _, option := range options {
		option(&pluginInfo)
	}

	app := app.New(app.UserOptions{}, allMocks.ToAppBoundaries())
	app.PluginService.Install(pluginInfo)
}

func TestInstall(t *testing.T) {
	t.Run("it should call metadata service", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		whenInstallPluginIsInvoked(allMocks)

		allMocks.Metadata.AssertCalled(t, "Get")
	})

	t.Run("it should check, if the plugin already exists", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		allMocks.Metadata.GivenGetReturns(metadata.Info{HomeDir: "/usr/some/"})

		whenInstallPluginIsInvoked(allMocks, pluginName("some-plugin"), version("1.0.1"))

		allMocks.FileSystem.AssertCalled(t, "Exists", "/usr/some/.sauron/plugins/some-plugin/1.0.1/some-plugin")
	})

	t.Run("it should start plugin download, when it does not exist", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		allMocks.Metadata.GivenGetReturns(metadata.Info{Arch: "amd64", Os: "darwin"})

		whenInstallPluginIsInvoked(allMocks, pluginName("some-plugin"), version("1.0.1"), url("http://repo/someplugin"))

		allMocks.HttpClient.AssertCalled(t, "Get", "http://repo/someplugin/darwin-amd64-1.0.1.tar.gz")
	})

	t.Run("it should save the downloaded archive in the cache", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		allMocks.Metadata.GivenGetReturns(metadata.Info{Arch: "amd64", Os: "darwin", HomeDir: "/usr/some/"})
		expectedBody := mocks.NewMockResponseBody()
		allMocks.HttpClient.GivenGetReturns(expectedBody)

		whenInstallPluginIsInvoked(allMocks, pluginName("some-plugin"), version("1.0.1"))

		allMocks.FileSystem.AssertCalled(t, "Save", "/usr/some/.sauron/cache/some-plugin/1.0.1/darwin-amd64-1.0.1.tar.gz", expectedBody)
	})

	t.Run("it should close data stream, after it was saved", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		expectedBody := mocks.NewMockResponseBody()
		allMocks.HttpClient.GivenGetReturns(expectedBody)

		whenInstallPluginIsInvoked(allMocks)

		expectedBody.AssertCalled(t, "Close")
	})
}
