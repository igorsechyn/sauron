package test

import (
	"github.com/igorsechyn/sauron/mocks"
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/cli"
)

func WhenCliCommandIsInvoked(allMocks *mocks.Mocks, userOptions app.UserOptions, args ...string) error {
	application := app.App{
		CliName:          userOptions.CliName,
		LongDescription:  userOptions.LongDescription,
		ShortDescription: userOptions.ShortDescription,
		Version:          userOptions.Version,
		ConsoleWriter:    allMocks.ConsoleWriter,
		Installer:        allMocks.PluginInstaller,
	}
	someCli := cli.NewWithApp(application)
	return someCli.RunWithArgs(args...)
}
