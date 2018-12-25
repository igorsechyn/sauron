package version_test

import (
	"testing"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/test"
	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	t.Run("it should include version command in available commands", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()

		test.WhenCliCommandIsInvoked(allMocks, app.UserOptions{CliName: "some-cli"}, "help")

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, "Prints version of some-cli")
	})

	t.Run("it should print the version, when version command is invoked", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		test.WhenCliCommandIsInvoked(
			allMocks,
			app.UserOptions{Version: "dev"},
			"version",
		)

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Equal(t, output, "dev")
	})
}
