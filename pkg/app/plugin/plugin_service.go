package plugin

import (
	"path"

	"github.com/igorsechyn/sauron/pkg/app/files"
	"github.com/igorsechyn/sauron/pkg/app/metadata"
)

type Info struct {
	PluginName       string
	Version          string
	URL              string
	ShortDescription string
	LongDescription  string
}

type Service interface {
	Install(pluginInfo Info) error
}

type service struct {
	metadata   metadata.Metadata
	fileSystem files.FileSystem
}

func NewService(metadata metadata.Metadata, fileSystem files.FileSystem) Service {
	return &service{
		metadata:   metadata,
		fileSystem: fileSystem,
	}
}

func (service *service) Install(info Info) error {
	meta := service.metadata.Get()
	path := path.Join(meta.HomeDir, ".sauron", "cache", meta.Os, meta.Arch, info.PluginName, info.Version, info.PluginName)
	service.fileSystem.Exists(path)
	return nil
}
