package http_test

import (
	"testing"

	"github.com/igorsechyn/sauron/pkg/app/metadata"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
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

func whenInstallAPluginIsInvoked(allMocks *mocks.Mocks, options ...installOption) {
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
	app.Installer.Install(pluginInfo)
}

func TestInstall(t *testing.T) {
	t.Run("it should call metadata service", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		whenInstallAPluginIsInvoked(allMocks)

		allMocks.Metadata.AssertCalled(t, "Get")
	})

	t.Run("it should check, if the plugin already exists", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		allMocks.Metadata.GivenGetReturns(metadata.Info{HomeDir: "/usr/some/", Arch: "amd64", Os: "darwin"})

		whenInstallAPluginIsInvoked(allMocks, pluginName("some-plugin"), version("1.0.1"))

		allMocks.FileSystem.AssertCalled(t, "Exists", "/usr/some/.sauron/cache/darwin/amd64/some-plugin/1.0.1/some-plugin")
	})
}
