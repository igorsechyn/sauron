package http

import (
	"path"

	"github.com/igorsechyn/sauron/pkg/app/files"
	"github.com/igorsechyn/sauron/pkg/app/metadata"
	"github.com/igorsechyn/sauron/pkg/app/plugin"
)

type PluginInstaller struct {
	Metadata   metadata.Metadata
	FileSystem files.FileSystem
}

func (installer *PluginInstaller) Install(options plugin.Info) error {
	meta := installer.Metadata.Get()
	path := path.Join(meta.HomeDir, ".sauron", "cache", meta.Os, meta.Arch, options.PluginName, options.Version, options.PluginName)
	installer.FileSystem.Exists(path)
	return nil
}
