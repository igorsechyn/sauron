package root

import (
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/cmd"
	scobra "github.com/igorsechyn/sauron/pkg/cmd/cobra"
	"github.com/spf13/cobra"
)

func NewCmd(app app.App) cmd.Command {
	cmd := &cobra.Command{
		Use:   app.CliName,
		Long:  app.LongDescription,
		Short: app.ShortDescription,
	}
	cmd.SetOutput(app.ConsoleWriter)
	return scobra.NewCommand(cmd)
}
