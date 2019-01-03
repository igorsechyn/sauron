package app

import (
	"io"

	"github.com/igorsechyn/sauron/pkg/app/http"

	"github.com/igorsechyn/sauron/pkg/app/files"

	"github.com/igorsechyn/sauron/pkg/app/metadata"

	"github.com/igorsechyn/sauron/pkg/app/plugin"
)

type Boundaries struct {
	ConsoleWriter io.Writer
	Metadata      metadata.Metadata
	FileSystem    files.FileSystem
	HttpClient    http.Client
}

type UserOptions struct {
	CliName          string
	ShortDescription string
	LongDescription  string
	Version          string
}

type App struct {
	ConsoleWriter    io.Writer
	PluginService    plugin.Service
	CliName          string
	ShortDescription string
	LongDescription  string
	Version          string
}

func New(config UserOptions, boundaries Boundaries) App {
	paths := &plugin.Paths{
		Metadata: boundaries.Metadata,
	}
	return App{
		ConsoleWriter:    boundaries.ConsoleWriter,
		CliName:          config.CliName,
		ShortDescription: config.ShortDescription,
		LongDescription:  config.LongDescription,
		Version:          config.Version,
		PluginService:    plugin.NewService(paths, boundaries.FileSystem, boundaries.HttpClient),
	}
}
