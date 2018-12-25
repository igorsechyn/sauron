package cli_test

import (
	"testing"

	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/test"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	t.Run("it should use provided cli config for usage output", func(t *testing.T) {
		allMocks := mocks.InitAllMocks()
		test.WhenCliCommandIsInvoked(
			allMocks,
			app.UserOptions{
				CliName: "some-cli", ShortDescription: "some-cli short", LongDescription: "some-cli long desc",
			},
		)

		output := allMocks.ConsoleWriter.GetOutput()
		assert.Contains(t, output, `Use "some-cli [command]`)
	})
}
