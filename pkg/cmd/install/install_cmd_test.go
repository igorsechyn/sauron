package install_test

import (
	"fmt"
	"testing"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/app/plugin"
	"github.com/igorsechyn/sauron/test"
	"github.com/stretchr/testify/assert"
)

func TestInstall_Command(t *testing.T) {
	t.Run("it should include install command", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "help")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, "install")
	})

	t.Run("it should print a short description for install command", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "help")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, "Installs additional plugins and records them in the local manifest")
	})

	t.Run("it should fail, if no arguments are supplied", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, "Error: accepts 1 arg(s), received 0")
	})

	t.Run("it should fail, if more than 2 arguments are supplied", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install", "first", "second")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, "Error: accepts 1 arg(s), received 2")
	})

	t.Run("it should fail, if plugin url location flag is missing", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install", "first", "--version", "1.0.1")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, `Error: required flag(s) "url" not set`)
	})

	t.Run("it should fail, if plugin version flag is missing", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install", "plugin-name", "--url", "http://some.url")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, `Error: required flag(s) "version" not set`)
	})

	t.Run("it should call plugin installer with options and default values for descriptions", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install", "plugin-name", "--url", "http://some.url", "--version", "0.1.0")

		allMocks.PluginInstaller.AssertCalled(t, "Install", plugin.Info{
			ShortDescription: "plugin-name", LongDescription: "plugin-name", PluginName: "plugin-name", URL: "http://some.url", Version: "0.1.0",
		})
	})

	t.Run("it should call plugin installer with options", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(
			allMocks,
			app.UserOptions{},
			"install", "plugin-name", "--url", "http://some.url", "--version", "0.1.0", "--short", "short description", "--long", "long description")

		allMocks.PluginInstaller.AssertCalled(t, "Install", plugin.Info{
			ShortDescription: "short description", LongDescription: "long description", PluginName: "plugin-name", URL: "http://some.url", Version: "0.1.0",
		})
	})

	t.Run("it should return an error, if installing plugin fails", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		allMocks.PluginInstaller.GivenInstallFails(fmt.Errorf("could not find plugin"))

		err := test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{}, "install", "plugin-name", "--url", "http://some.url", "--version", "0.1.0")

		assert.EqualError(t, err, "could not find plugin")
	})
}
