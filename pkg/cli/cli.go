package cli

import (
	"os"

	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/cmd"
	"github.com/igorsechyn/sauron/pkg/cmd/install"
	"github.com/igorsechyn/sauron/pkg/cmd/root"
	"github.com/igorsechyn/sauron/pkg/cmd/version"
)

func New(userOptions app.UserOptions) *Cli {
	productionBoundaries := app.Boundaries{
		ConsoleWriter: os.Stdout,
	}
	return NewWithBoundaries(userOptions, productionBoundaries)
}

func NewWithBoundaries(userOptions app.UserOptions, boundaries app.Boundaries) *Cli {
	application := app.New(userOptions, boundaries)

	return NewWithApp(application)
}

func NewWithApp(application app.App) *Cli {
	rootCmd := root.NewCmd(application)
	versionCmd := version.NewCmd(application)
	installCmd := install.NewCmd(application)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(installCmd)
	return &Cli{rootCmd: rootCmd}
}

type Cli struct {
	rootCmd cmd.Command
}

func (cli *Cli) RunWithArgs(args ...string) error {
	return cli.rootCmd.RunWithArgs(args...)
}
