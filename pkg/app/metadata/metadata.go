package metadata

type Info struct {
	HomeDir string
	Os      string
	Arch    string
}

type Metadata interface {
	Get() Info
}
