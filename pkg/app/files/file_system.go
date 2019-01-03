package files

import (
	"io"
)

type FileSystem interface {
	Exists(path string) bool
	Save(path string, data io.Reader) error
}
