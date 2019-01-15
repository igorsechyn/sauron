# Sauron
> Simple and minimilistic CLI framework to compose already existing CLIs into a homogeneuous experience. Existing CLIs must be provided for download as executables in a specific format

## Requirements
- go 1.11 or higher

## Installation

Get the `sauron` as a dependency

```
go get github.com/igorsechyn/httptest-interaction-listener
```

## Design
We often find work environments with multiple CLIs provided by different teams for different purposes. Different CLIs may use different package managers and may require specific runtime environments. Creating a monolithic CLI would put too much burden on dev experience. `Sauron`'s idea is to provide a framework for a generic CLI, which allows composing existing CLIs under one umbrella. The only requirement is to package existing CLIs as executables and provide them for download.

Out of the box the CLI provides only two commands `version` and `install`. 

`install` command takes the name of a plugin, URL location, where executable can be downloaded, and the version (see further for repository layout of executables). After successfully downloading the executable and putting it in `~/.sauron/cache/<plugin-name>/<version>`, the framework will register additional command under `<plugin-name>` and delegate calls to the downloaded executable. It will also write a local manifest file with a list of installed plugins. This manifest file will only be loaded, when the cli is run from that directory. It is also possible to provide a global manifest file under `HOME` directory or when compiling the CLI

### Manifest
Manifest file contains information about available plugins. Manifest can be provided in three different ways:

1. As a parameter to initialise a CLI during compile time
2. As a global manifest in a file located under `~/.sauron/manifest.json`
3. As a local manifest file `manifest.json` in the working directory, where the CLI is executed

The list of available plugins is built from all available manifests in the order, as they listed above. I.e. if the global manifest contains a plugin with version 1.0.0 and a local manifest has the version 2.0.0, version 2.0.0 will be used

A manifest file has the following form:

```
{
    "plugins": [
        {
            "name": "plugin-name",
            "version": "1.0.1",
            "url": "https://repo.org/location",
            "short": "Optional short description (defaults to name)",
            "long": "Optional long description (defaults to name)"
        }
    ]
}
```

### Plugins
A plugin is any executable, that can be invoked from the command line. In order to be installed, the executble must be provided for download and named according to the following schema:

```
<os>-<arch>-<version>
```
For example `darwin-amd64-1.0.0`. At runtime the cli will construct the name based on the options and the platform it is being executed on.

## Usage


## Changelog
See [CHANGELOG.md](CHANGELOG.md)

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md)

## License
See [LICENSE.txt](LICENSE.txt)