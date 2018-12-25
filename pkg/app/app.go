package app

import (
	"io"

	"github.com/igorsechyn/sauron/pkg/app/files"

	"github.com/igorsechyn/sauron/pkg/app/metadata"

	"github.com/igorsechyn/sauron/pkg/app/plugin"
	"github.com/igorsechyn/sauron/pkg/app/plugin/http"
)

type Boundaries struct {
	ConsoleWriter io.Writer
	Metadata      metadata.Metadata
	FileSystem    files.FileSystem
}

type UserOptions struct {
	CliName          string
	ShortDescription string
	LongDescription  string
	Version          string
}

type App struct {
	ConsoleWriter    io.Writer
	Installer        plugin.Installer
	CliName          string
	ShortDescription string
	LongDescription  string
	Version          string
}

func New(config UserOptions, boundaries Boundaries) App {
	return App{
		ConsoleWriter:    boundaries.ConsoleWriter,
		CliName:          config.CliName,
		ShortDescription: config.ShortDescription,
		LongDescription:  config.LongDescription,
		Version:          config.Version,
		Installer:        &http.PluginInstaller{Metadata: boundaries.Metadata, FileSystem: boundaries.FileSystem},
	}
}
