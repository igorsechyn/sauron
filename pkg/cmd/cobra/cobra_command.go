package cobra

import (
	"github.com/igorsechyn/sauron/pkg/cmd"
	"github.com/spf13/cobra"
)

func NewCommand(cmd *cobra.Command) cmd.Command {
	return &Command{Cmd: cmd}
}

type Command struct {
	Cmd *cobra.Command
}

func (cmd *Command) RunWithArgs(args ...string) error {
	cmd.Cmd.SetArgs(args)
	return cmd.Cmd.Execute()
}

func (cmd *Command) AddCommand(subCommand cmd.Command) {
	cobraSubCommand, ok := subCommand.(*Command)
	if ok {
		cmd.Cmd.AddCommand(cobraSubCommand.Cmd)
	}
}
