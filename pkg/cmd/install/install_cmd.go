package install

import (
	"github.com/igorsechyn/sauron/pkg/app"
	"github.com/igorsechyn/sauron/pkg/app/plugin"
	"github.com/igorsechyn/sauron/pkg/cmd"
	scobra "github.com/igorsechyn/sauron/pkg/cmd/cobra"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func NewCmd(application app.App) cmd.Command {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Installs additional plugins and records them in the local manifest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			installOptions, err := getInstallOptions(cmd.Flags(), args[0])
			if err != nil {
				return err
			}
			return application.Installer.Install(installOptions)
		},
	}

	installCmd.Flags().String("url", "", "Download URL for a plugin (without the name of the executable")
	installCmd.Flags().String("version", "", "Semantic version of the plugin")
	installCmd.Flags().String("short", "", "Commands short description")
	installCmd.Flags().String("long", "", "Commands long description")
	err := installCmd.MarkFlagRequired("url")
	if err != nil {
		return nil
	}
	err = installCmd.MarkFlagRequired("version")
	if err != nil {
		return nil
	}
	return scobra.NewCommand(installCmd)
}

func getInstallOptions(flags *flag.FlagSet, pluginName string) (plugin.Info, error) {
	url, err := flags.GetString("url")
	if err != nil {
		return plugin.Info{}, err
	}
	version, err := flags.GetString("version")
	if err != nil {
		return plugin.Info{}, err
	}
	shortDescription := getFlagOrDefault(flags, "short", pluginName)
	longDescription := getFlagOrDefault(flags, "long", pluginName)

	return plugin.Info{
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
		PluginName:       pluginName,
		Version:          version,
		URL:              url,
	}, nil
}

func getFlagOrDefault(flags *flag.FlagSet, flagName, defaultValue string) string {
	value, err := flags.GetString(flagName)
	if err != nil || value == "" {
		return defaultValue
	}
	return value
}
