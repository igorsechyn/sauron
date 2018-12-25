package plugin

type InstallOptions struct {
	PluginName       string
	Version          string
	URL              string
	ShortDescription string
	LongDescription  string
}

type Installer interface {
	Install(options InstallOptions) error
}
