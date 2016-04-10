package backends

import (
	"errors"

	"github.com/jacoelho/flux/backends/kafka"
	"github.com/jacoelho/flux/backends/stdout"
)

type Client interface {
	Write(p []byte) (n int, err error)
	Flush() error
	Close() error
}

var ErrInvalidBackend = errors.New("invalid backend")

func NewClient(c *Config) (Client, error) {
	switch c.Type {
	case "kafka":
		return kafka.New(c.Nodes)
	case "stdout":
		return stdout.New()
	}
	return nil, ErrInvalidBackend
}
