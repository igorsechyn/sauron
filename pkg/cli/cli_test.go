package cli_test

import (
	"testing"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/test"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	t.Run("it should use provided cli name for usage output", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		test.WhenCliCommandIsInvoked(
			allMocks,
			app.UserOptions{
				CliName: "some-cli",
			},
		)

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, `Use "some-cli [command]`)
	})

	t.Run("it should use provided cli description for usage output", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		test.WhenCliCommandIsInvoked(
			allMocks,
			app.UserOptions{
				CliName: "some-cli", ShortDescription: "short description",
			},
		)

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, `short description`)
	})
}
