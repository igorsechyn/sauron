package files

type FileSystem interface {
	Exists(path string) bool
}
