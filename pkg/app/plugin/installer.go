package plugin

type Info struct {
	PluginName       string
	Version          string
	URL              string
	ShortDescription string
	LongDescription  string
}

type Installer interface {
	Install(pluginInfo Info) error
}
