package plugin

import (
	"github.com/igorsechyn/sauron/pkg/app/http"

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
	httpClient http.Client
	paths      *Paths
}

func NewService(paths *Paths, fileSystem files.FileSystem, httpClient http.Client) Service {
	return &service{
		fileSystem: fileSystem,
		httpClient: httpClient,
		paths:      paths,
	}
}

func (service *service) Install(info Info) error {
	pluginPath := service.paths.GetPluginPath(info)
	service.fileSystem.Exists(pluginPath)
	data, _ := service.httpClient.Get(info.URL + "/" + service.paths.GetArchiveName(info))
	defer data.Close()
	pluginCachePath := service.paths.GetCachePath(info)
	service.fileSystem.Save(pluginCachePath, data)
	return nil
}
