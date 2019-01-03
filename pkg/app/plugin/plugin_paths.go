package plugin

import (
	"fmt"
	"path"

	"github.com/igorsechyn/sauron/pkg/app/metadata"
)

const (
	sauronDir  = ".sauron"
	pluginsDir = "plugins"
	cacheDir   = "cache"
)

type Paths struct {
	Metadata metadata.Metadata
}

func (paths *Paths) GetPluginPath(info Info) string {
	meta := paths.Metadata.Get()
	return path.Join(meta.HomeDir, sauronDir, pluginsDir, info.PluginName, info.Version, info.PluginName)
}

func (paths *Paths) GetCachePath(info Info) string {
	meta := paths.Metadata.Get()
	return path.Join(meta.HomeDir, sauronDir, cacheDir, info.PluginName, info.Version, paths.GetArchiveName(info))
}

func (paths *Paths) GetArchiveName(info Info) string {
	meta := paths.Metadata.Get()
	return fmt.Sprintf("%v-%v-%v.tar.gz", meta.Os, meta.Arch, info.Version)
}
