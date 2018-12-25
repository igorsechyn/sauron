package metadata

type Info struct{}

type Metadata interface {
	Get() Info
}
