package http

import "io"

type Client interface {
	Get(url string) (io.ReadCloser, error)
}
