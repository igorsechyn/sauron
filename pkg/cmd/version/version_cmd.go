package version

import (
	"fmt"

	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/cmd"
	scobra "github.com/igorsechyn/sauron/pkg/cmd/cobra"
	"github.com/spf13/cobra"
)

func NewCmd(app app.App) cmd.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Prints version of %v", app.CliName),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(app.Version)
		},
	}

	return scobra.NewCommand(versionCmd)
}
