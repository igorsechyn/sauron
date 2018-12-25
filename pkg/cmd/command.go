package cmd

type Command interface {
	RunWithArgs(args ...string) error
	AddCommand(cmd Command)
}
